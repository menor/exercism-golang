package luhn

import (
	"fmt"
	"strings"
	"unicode"
)

func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	if len(s) <= 1 {
		return false
	}

	var num int
	double := "even"

	if len(s)%2 == 0 {
		double = "odd"
	}

	for i, char := range s {
		if !unicode.IsDigit(char) {
			fmt.Println(s[i])
			return false
		}

		value := int(char - '0')

		if (double == "even" && (i+1)%2 == 0) || (double == "odd" && (i+1)%2 != 0) {
			value = value * 2

			if value > 9 {
				value = value - 9
			}

		}

		num += value
		fmt.Println(i, value)
	}

	fmt.Println(num)
	return num%10 == 0
}

// 1234567 -> 642
// 123456 -> 531
