package isogram

import "strings"

func IsIsogram(s string) bool {
	s = strings.ToLower(s)

	letters := make(map[rune]bool, len(s))

	for _, letter := range s {
		if letters[letter] {
			return false
		}

		if !(letter == ' ' || letter == '-') {
			letters[letter] = true
		}
	}

	return true
}
