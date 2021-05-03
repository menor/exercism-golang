package grains

import "fmt"

// Square returns the amount of grains for a given position in the board
func Square(n int) (uint64, error) {
	var result uint64 = 1

	if n < 1 || n > 64 {
		return 0, fmt.Errorf("number must be greater than 0 and smaller than 65, go %v", n)
	}

	// we use the left shift operator to get the square of 2 to the number
	// this is the same as math.Pow(2, float64(n-1) but faster
	result = 1 << (n - 1)

	return result, nil
}

// Total returns the total number of grains, the board contains
func Total() uint64 {
	// We already know the total so no need to
	// calculate it every time
	const total uint64 = 18446744073709551615

	return total
}
