package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes several texts and returns a
// combined count of the letter frequency
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap, len(texts))
	result := FreqMap{}

	for _, text := range texts {
		// We need to use a function literal here, so the text is passed
		// by value and each Frequency invocation gets its own copy of the
		// text, this avoids every goroutine just getting the first value
		go func(t string) {
			c <- Frequency(t)
		}(text)
	}

	// since we started len(texts) goroutines, we need to wait
	// for len(texts) goroutines to end, this works cause we keep
	// the program running (blocked) until we receive the number of
	// responses that we started
	for range texts {
		mapped := <-c

		for k, v := range mapped {
			result[k] += v
		}
	}

	return result
}
