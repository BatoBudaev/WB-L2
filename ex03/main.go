package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	kFlag = flag.Int("k", -1, "Указание колонки для сортировки")
	nFlag = flag.Bool("n", false, "Сортировать по числовому значению")
	rFlag = flag.Bool("r", false, "Сортировать в обратном порядке")
	uFlag = flag.Bool("u", false, "Не выводить повторяющиеся строки")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <inputfile> [-k <column>] [-n] [-r] [-u]")
		os.Exit(1)
	}

	flag.Parse()

	inputFile := os.Args[len(os.Args)-1]
	sortLines(inputFile, *kFlag, *nFlag, *rFlag, *uFlag)
}

func sortLines(inputFile string, kFlag int, nFlag, rFlag, uFlag bool) {
	// Открытие файла для чтения
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %v\n", inputFile, err)
		return
	}
	defer file.Close()

	// Создание сканера для чтения файла
	scanner := bufio.NewScanner(file)
	var lines []string // Инициализация слайса для хранения строк

	// Чтение содержимого файла построчно
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Проверка на наличие ошибок при чтении файла
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения файла %s: %v\n", inputFile, err)
		return
	}

	// Вызов функции для сортировки строк
	lines = sortLinesFunc(lines, kFlag, nFlag, rFlag, uFlag)

	// Формирование имени файла для записи отсортированных строк
	outputFileName := "sorted_" + inputFile

	// Создание файла для записи отсортированных строк
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Ошибка создания файла %s: %v\n", outputFileName, err)
		return
	}
	defer outputFile.Close()

	// Создание буферизованного писателя для записи в файл
	writer := bufio.NewWriter(outputFile)

	// Запись отсортированных строк в файл
	for _, line := range lines {
		fmt.Fprintln(writer, line) // Запись строки в файл
	}
	writer.Flush() // Очистка буфера и запись оставшихся данных в файл
}

func sortLinesFunc(lines []string, kFlag int, nFlag, rFlag, uFlag bool) []string {
	// Проверка на пустоту списка строк
	if len(lines) == 0 {
		return lines // Ранний выход, если список пуст
	}

	// Сортировка списка строк
	sort.SliceStable(lines, func(i, j int) bool {
		// Если установлен флаг -n, сортировка по числовому значению
		if nFlag {
			// Получение числового значения из указанной колонки
			numI, err := strconv.Atoi(getColumn(lines[i], kFlag))
			if err != nil {
				return false
			}
			numJ, err := strconv.Atoi(getColumn(lines[j], kFlag))
			if err != nil {
				return false
			}
			// Сравнение числовых значений
			return numI < numJ
		} else {
			// Получение индекса колонки для сортировки
			colIndex := getColIndex(kFlag)
			// Проверка, находится ли индекс колонки в допустимом диапазоне
			if colIndex < 0 || colIndex >= len(strings.Fields(lines[i])) {
				return false
			}
			// Сортировка по алфавиту, используя значения из указанной колонки
			lineI := strings.Fields(lines[i])[colIndex]
			lineJ := strings.Fields(lines[j])[colIndex]
			return lineI < lineJ
		}
	})

	// Обратная сортировка, если установлен флаг -r
	if rFlag {
		sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	}

	// Удаление дубликатов, если установлен флаг -u
	if uFlag {
		lines = removeDuplicates(lines)
	}

	return lines
}

func getColumn(line string, col int) string {
	parts := strings.Fields(line)
	if col >= 0 && col < len(parts) {
		return parts[col]
	}
	return ""
}

func getColIndex(col int) int {
	return col - 1 // Индексы начинаются с 0, поэтому вычитаем 1
}

func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, line := range lines {
		normLine := strings.TrimSpace(strings.ToLower(line))
		if !seen[normLine] {
			seen[normLine] = true
			result = append(result, line)
		}
	}

	return result
}
