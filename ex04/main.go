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
		sort.Slice(runeSlice, func(i, j int) bool { return runeSlice[i] < runeSlice[j] })
		sortedWord := string(runeSlice)
		if _, exists := anagrams[sortedWord]; !exists {
			anagrams[sortedWord] = []string{loweredWord}
		} else {
			anagrams[sortedWord] = append(anagrams[sortedWord], loweredWord)
		}
	}

	// Фильтрация групп, оставляя только те, которые содержат более одного слова
	result := make(map[string][]string)
	for sortedWord, group := range anagrams {
		if len(group) > 1 {
			result[sortedWord] = group
		}
	}

	return result
}

// Сортирует символы строки в алфавитном порядке
func sortWords(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	groups := findAnagrams(words)
	fmt.Printf("%v\n", groups)
}
