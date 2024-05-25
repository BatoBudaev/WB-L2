package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	aFlag     = flag.Int("A", 0, "Печатать +N строк после совпадения")
	bFlag     = flag.Int("B", 0, "Печатать +N строк до совпадения")
	cFlag     = flag.Int("C", 0, "Печатать ±N строк вокруг совпадения")
	countFlag = flag.Bool("c", false, "Количество строк")
	iFlag     = flag.Bool("i", false, "Игнорировать регистр")
	vFlag     = flag.Bool("v", false, "Инвертировать фильтр (исключать совпадения)")
	fixedFlag = flag.Bool("F", false, "Точное совпадение со строкой, не паттерн")
	nFlag     = flag.Bool("n", false, "Печатать номер строки")
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Использование: go run main.go [OPTIONS] PATTERN FILE")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	fileName := flag.Arg(1)

	grepFile(pattern, fileName, *aFlag, *bFlag, *cFlag, *countFlag, *iFlag, *vFlag, *fixedFlag, *nFlag)
}

func grepFile(pattern, fileName string, aFlag, bFlag, cFlag int, countFlag, iFlag, vFlag, fixedFlag, nFlag bool) {
	// Открытие файла для чтения
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Обработка флага игнорирования регистра
	if iFlag && !fixedFlag {
		pattern = "(?i)" + pattern
	}

	var re *regexp.Regexp
	if !fixedFlag {
		// Компиляция регулярного выражения
		re, err = regexp.Compile(pattern)
		if err != nil {
			fmt.Printf("Ошибка компиляции регулярного выражения: %v\n", err)
			return
		}
	}

	// Создание сканера для чтения файла
	scanner := bufio.NewScanner(file)
	var lines []string // Инициализация слайса для хранения строк

	// Чтение содержимого файла построчно
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Проверка на наличие ошибок при чтении файла
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения файла %s: %v\n", fileName, err)
		return
	}

	// Обработка контекста
	if cFlag > 0 {
		aFlag, bFlag = cFlag, cFlag
	}

	matchedLines := make(map[int]bool)
	for i, line := range lines {
		matched := false
		if fixedFlag {
			matched = line == pattern
		} else {
			matched = re.MatchString(line)
		}
		if vFlag {
			matched = !matched
		}
		if matched {
			matchedLines[i] = true
			for j := 1; j <= aFlag; j++ {
				if i+j < len(lines) {
					matchedLines[i+j] = true
				}
			}
			for j := 1; j <= bFlag; j++ {
				if i-j >= 0 {
					matchedLines[i-j] = true
				}
			}
		}
	}

	if countFlag {
		fmt.Println(len(matchedLines))
		return
	}

	for i, line := range lines {
		if matchedLines[i] {
			if nFlag {
				fmt.Printf("%d:%s\n", i+1, line)
			} else {
				fmt.Println(line)
			}
		}
	}
}
