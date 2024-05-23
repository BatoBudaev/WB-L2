package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		loweredWord := strings.ToLower(word) // В нижний регистр
		runeSlice := []rune(loweredWord)
		sort.Slice(runeSlice, func(i, j int) bool { return runeSlice[i] < runeSlice[j] }) // Сортировка символов строки в алфавитном порядке
		sortedWord := string(runeSlice)
		if _, exists := anagrams[sortedWord]; !exists {
			anagrams[sortedWord] = []string{loweredWord}
		} else {
			anagrams[sortedWord] = append(anagrams[sortedWord], loweredWord)
		}
	}

	// Создаем карту результатов, где ключ - это первое слово из группы анаграмм
	result := make(map[string][]string)
	for _, group := range anagrams {
		firstWord := group[0]     // Берем первое слово из группы
		result[firstWord] = group // Используем его как ключ
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	groups := findAnagrams(words)
	fmt.Printf("%v\n", groups)
}
