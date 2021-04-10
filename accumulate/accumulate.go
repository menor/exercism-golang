package accumulate

type apply func(string) string

func Accumulate(src []string, fn apply) []string {
	var result = make([]string, 0, len(src))

	for _, a := range src {
		result = append(result, fn(a))
	}

	return result
}
