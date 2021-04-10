package etl

import (
	"strings"
)

type Legacy map[int][]string
type Transformed map[string]int

func Transform(legacy Legacy) Transformed {
	transformed := make(Transformed)

	for value, letters := range legacy {
		for _, letter := range letters {
			transformed[strings.ToLower(letter)] = value
		}
	}

	return transformed
}
