package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if a string is valid per the Luhn formula
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	if len(s) <= 1 {
		return false
	}

	var num int

	shouldDouble := len(s)%2 == 0

	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}

		value := int(char - '0')

		if shouldDouble {
			value *= 2

			if value > 9 {
				value -= 9
			}
		}

		num += value
		shouldDouble = !shouldDouble
	}

	return num%10 == 0
}
