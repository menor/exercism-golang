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

	shouldDouble := length%2 == 0

	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}

		value := int(char - '0')

		if shouldDouble {
			value = doubleAndReduce(value)
		}

		num += value
		shouldDouble = !shouldDouble
	}

	return num%10 == 0
}

func doubleAndReduce(n int) int {
	n *= 2

	if n > 9 {
		n -= 9
	}

	return n
}
