package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

// WordCount takes a string and counts the occurrences of
// words in it
func WordCount(s string) Frequency {
	count := make(Frequency)

	words := extractWords(s)

	for _, word := range words {
		word = strings.ToLower(word)
		word = trimWord(word)

		count[word]++
	}

	return count
}

// extractWords takes a string and gives you a slice with every word and
// numbers, it respects apostrophes, so it will leave quotes around words
func extractWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})
}

// trimWords removes spaces and quotes from around a passed string
func trimWord(s string) string {
	s = strings.Trim(s, "")
	s = strings.TrimPrefix(s, "'")
	s = strings.TrimRight(s, "'")

	return s
}
