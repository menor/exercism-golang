package anagram

import (
	"reflect"
	"strings"
)

func Detect(s string, candidates []string) []string {
	word := countLetters(s)
	var matches []string

	for _, candidate := range candidates {

		if strings.EqualFold(s, candidate) {
			continue
		}

		letters := countLetters(candidate)

		if reflect.DeepEqual(word, letters) {
			matches = append(matches, candidate)
		}

	}

	return matches
}

func countLetters(s string) map[rune]int {
	s = strings.ToLower(s)
	counter := make(map[rune]int)

	for _, r := range s {
		counter[r]++
	}

	return counter
}
