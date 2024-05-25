package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	fields    = flag.String("f", "", "Выбрать поля (колонки)")
	delimiter = flag.String("d", "\t", "Использовать другой разделитель")
	separated = flag.Bool("s", false, "Только строки с разделителем")
)

func main() {
	flag.Parse()

	if *fields == "" {
		fmt.Println("Использование: go run main.go -f FIELDS [-d DELIMITER] [-s]")
		os.Exit(1)
	}

	// Парсинг полей
	fieldIndexes, err := parseFields(*fields)
	if err != nil {
		fmt.Printf("Ошибка парсинга полей: %v\n", err)
		os.Exit(1)
	}

	// Чтение из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// Разделение строки по заданному разделителю
		columns := strings.Split(line, *delimiter)

		// Если установлен флаг -s, пропускаем строки без разделителя
		if *separated && len(columns) < 2 {
			continue
		}

		var selectedColumns []string
		for _, index := range fieldIndexes {
			if index < len(columns) {
				selectedColumns = append(selectedColumns, columns[index])
			}
		}

		fmt.Println(strings.Join(selectedColumns, *delimiter))
	}

	// Проверка на наличие ошибок при чтении
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения из STDIN: %v\n", err)
	}
}

func parseFields(fields string) ([]int, error) {
	parts := strings.Split(fields, ",")
	var fieldIndexes []int
	for _, part := range parts {
		var index int
		_, err := fmt.Sscanf(part, "%d", &index)
		if err != nil {
			return nil, err
		}
		// Внутреннее представление индексов колонок как 0-индексируемых
		fieldIndexes = append(fieldIndexes, index-1)
	}
	return fieldIndexes, nil
}
