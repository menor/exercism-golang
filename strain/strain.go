package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(fn func(int) bool) Ints {
	var keep Ints

	for _, el := range i {
		if fn(el) {
			keep = append(keep, el)
		}
	}
	return keep
}

func (i Ints) Discard(fn func(int) bool) Ints {
	var keep Ints

	for _, el := range i {
		if !fn(el) {
			keep = append(keep, el)
		}
	}
	return keep
}

func (s Strings) Keep(fn func(string) bool) Strings {
	var keep Strings

	for _, el := range s {
		if fn(el) {
			keep = append(keep, el)
		}
	}
	return keep
}

func (l Lists) Keep(fn func([]int) bool) Lists {
	var keep Lists

	for _, el := range l {
		if fn(el) {
			keep = append(keep, el)
		}
	}
	return keep
}
