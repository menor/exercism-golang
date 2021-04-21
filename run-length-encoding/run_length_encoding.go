package encode

import (
	"strconv"
)

func RunLengthEncode(s string) string {
	if len(s) < 1 {
		return s
	}

	var result string
	runes := []rune(s)
	previous := runes[0]
	count := 1

	for _, l := range runes {
		if l == previous {
			count++
		} else {
			if count < 2 {
				result = result + string(l)
			} else {
				result = result + strconv.Itoa(count) + string(l)
			}
			count = 0
			previous = l
		}

	}
	return result
}

func RunLengthDecode(s string) string {
	return s
}
