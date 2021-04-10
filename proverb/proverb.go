// Package proverb generates a wise proverb
package proverb

import "fmt"

// Proverb generates a relevant proverb given a list of inputs
func Proverb(rhyme []string) []string {
	var proverb []string

	if len(rhyme) <= 0 {
		return proverb
	}

	for i, word := range rhyme {
		var line string

		if i == len(rhyme)-1 {
			line = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			line = fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[i+1])
		}

		proverb = append(proverb, line)
	}
	return proverb
}
