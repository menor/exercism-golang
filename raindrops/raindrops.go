package raindrops

import "strconv"

// Convert takes a number and depending on its factors
// returns a string representing raindrop sounds
func Convert(n int) string {
	result := ""

	if n%3 == 0 {
		result += "Pling"
	}

	if n%5 == 0 {
		result += "Plang"
	}

	if n%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(n)
	}

	return result
}
