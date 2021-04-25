package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int

	for i := 1; i < limit; i++ {

		for _, j := range divisors {

			if j == 0 {
				break
			}

			if i%j == 0 {
				sum += i
				break
			}

		}
	}

	return sum
}
