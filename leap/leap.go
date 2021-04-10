// Package leap provides utils to deal with leap years
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear takes a year and returns true if it's a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
