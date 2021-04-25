package grains

import (
	"fmt"
	"math"
)

func Square(n int) (uint64, error) {
	var result uint64 = 1

	if n < 1 || n > 64 {
		return 0, fmt.Errorf("number must be greater than 0 and smaller than 65, go %v", n)
	}

	result = uint64(math.Pow(2, float64(n-1)))

	return result, nil
}

func Total() uint64 {
	const total uint64 = 18446744073709551615

	return total
}
