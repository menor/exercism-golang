// Package twofer implements a solution for an exercism exercise
// When a name is given it returns the string "One for <name>, one for me"
// If no name is given it returns "One for you, one for me"
package twofer

// ShareWith takes a name and shares two thing with the name
// If no name is given it will default to you
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
