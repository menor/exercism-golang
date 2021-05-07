package clock

import "fmt"

// Clock represents a clock with its hours and minutes as int
type Clock struct {
	hours   int
	minutes int
}

// New takes hours and minutes and returns a Clock
func New(h, m int) Clock {
	h, m = format(h, m)

	c := Clock{
		hours:   h,
		minutes: m,
	}

	return c
}

// String returns a string representation of a clock formatted as hh:mm
func (c Clock) String() string {
	h := fmt.Sprintf("%02d", c.hours)
	m := fmt.Sprintf("%02d", c.minutes)

	return h + ":" + m
}

// Add takes a number of minuts and adds them to a Clock time
func (c Clock) Add(minutes int) Clock {
	m := c.minutes + minutes

	hours, mins := format(c.hours, m)

	c.hours = hours
	c.minutes = mins

	return c
}

// Subtract takes a number of minuts and subtracts them from a Clock time
func (c Clock) Subtract(minutes int) Clock {
	m := c.minutes - minutes%60
	h := c.hours - minutes/60

	hours, mins := format(h, m)

	c.hours = hours
	c.minutes = mins

	return c
}

func format(h, m int) (hours, minutes int) {
	if m >= 60 || m <= -60 {
		h = h + m/60
		m = m % 60
	}

	if m < 0 {
		h = h - (1 + m/60)
		m = 60 + m
	}

	if h >= 24 || h <= -24 {
		h = h % 24
	}

	if h < 0 {
		h = 24 + h%24
	}

	return h, m
}
