package scrabble

import "unicode"

func Score(word string) int {
	score := 0

	for _, letter := range word {
		score += getLetterScore(letter)
	}

	return score
}

// func Score(word string) int {
// scores := generateScores()
// score := 0
//
// for _, letter := range word {
// score += scores[unicode.ToUpper(letter)]
// }
//
// return score
// }

func getLetterScore(letter rune) int {
	switch unicode.ToUpper(letter) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

// func generateScores() map[rune]int {
// 	scores := make(map[rune]int)
// 	scores['A'] = 1
// 	scores['E'] = 1
// 	scores['I'] = 1
// 	scores['O'] = 1
// 	scores['U'] = 1
// 	scores['L'] = 1
// 	scores['N'] = 1
// 	scores['R'] = 1
// 	scores['S'] = 1
// 	scores['T'] = 1
// 	scores['D'] = 2
// 	scores['G'] = 2
// 	scores['B'] = 3
// 	scores['C'] = 3
// 	scores['M'] = 3
// 	scores['P'] = 3
// 	scores['F'] = 4
// 	scores['H'] = 4
// 	scores['V'] = 4
// 	scores['W'] = 4
// 	scores['Y'] = 4
// 	scores['K'] = 5
// 	scores['J'] = 8
// 	scores['X'] = 8
// 	scores['Q'] = 10
// 	scores['Z'] = 10

// 	return scores
// }
