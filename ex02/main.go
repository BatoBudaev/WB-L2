package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Unpack(s string) (string, error) {
	var sb strings.Builder
	runes := []rune(s)
	i := 0

	for i < len(runes) {
		r := runes[i]

		if r == '\\' { // Обработка escape-последовательности
			if i+1 >= len(runes) {
				return "", errors.New("некорректная escape-последовательность")
			}
			sb.WriteRune(runes[i+1])
			i += 2
		} else if r >= '0' && r <= '9' { // Проверка на число после символа
			if i > 0 && !strings.ContainsRune("0123456789", runes[i-1]) {
				count, err := strconv.Atoi(string(r))
				if err != nil || count == 0 {
					return "", errors.New("некорректная строка")
				}
				sb.WriteString(strings.Repeat(string(runes[i-1]), count-1))
				i++
			} else {
				return "", errors.New("некорректная строка")
			}
		} else { // Обычный символ
			sb.WriteRune(r)
			i++
		}
	}

	return sb.String(), nil
}

func main() {
	tests := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	for _, test := range tests {
		got, err := Unpack(test)
		if err == nil {
			fmt.Printf("Unpack(%q) = %q\n", test, got)
		} else {
			fmt.Println(err)
		}
	}
}
