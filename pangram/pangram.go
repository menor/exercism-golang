package pangram

import "strings"

func IsPangram(s string) bool {
	var counter = generateAlphabetCounter()

	s = strings.ToLower(s)

	for _, r := range s {
		counter[r]++
	}

	for _, v := range counter {
		if v == 0 {
			return false
		}
	}

	return true
}

func generateAlphabetCounter() map[rune]int {
	var alphabet = map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}

	return alphabet
}
