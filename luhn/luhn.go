package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if a string is valid per the Luhn formula
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	length := len(s)

	if length <= 1 {
		return false
	}

	var num int

	for i, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}

		value := int(char - '0')

		if shouldDouble(length, i) {
			value = doubleAndReduce(value)
		}

		num += value
	}

	return num%10 == 0
}

func doubleAndReduce(n int) int {
	n = n * 2

	if n > 9 {
		n = n - 9
	}

	return n
}

func shouldDouble(length, position int) bool {
	return length%2 == position%2
}
