package hamming

import "errors"

func Distance(a, b string) (int, error) {
	runeA := []rune(a)
	runeB := []rune(b)
	distance := 0

	if len(runeA) != len(runeB) {
		return distance, errors.New("values must be of the same size")
	}

	for index, runeValue := range runeA {
		if runeB[index] != runeValue {
			distance++
		}
	}

	return distance, nil
}
