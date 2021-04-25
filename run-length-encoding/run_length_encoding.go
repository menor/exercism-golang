package encode

import (
	"strconv"
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
	return s
}
