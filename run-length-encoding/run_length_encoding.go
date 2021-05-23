package encode

import (
	"strconv"
	"unicode"
)

// AABBBCCCC

func RunLengthEncode(s string) string {
	if len(s) < 1 {
		return s
	}

	var result string
	var previous rune
	count := 0

	for i, l := range s {

		if i+1 == len(s) {
			count++
		}

		if i != 0 && (l != previous || i+1 == len(s)) {

			if count < 2 {
				result = result + string(previous)
			} else {
				result = result + strconv.Itoa(count) + string(previous)
			}

			previous = l
			count = 1

		} else {

			count++
			previous = l

		}

	}
	return result
}

func RunLengthDecode(s string) string {
	if len(s) < 1 {
		return s
	}

	var result string
	count := "0"

	for _, c := range s {
		if unicode.IsDigit(c) {
			count += string(c)
		} else {
			if count == "0" {
				result += string(c)
			} else {
				repeats, _ := strconv.Atoi(count)

				for i := 0; i < repeats; i++ {
					result += string(c)
				}

				count = "0"
			}

		}
	}

	return result
}
