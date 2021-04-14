package reverse

import (
	"strings"
)

func Reverse(s string) string {
	var reversed strings.Builder

	r := []rune(s)
	l := len(r)

	for i := l - 1; i >= 0; i-- {
		reversed.WriteRune(r[i])
	}

	return reversed.String()
}
