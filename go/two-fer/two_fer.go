// Package twofer implements a function for sharing.
package twofer

// ShareWith helps you to share another person.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
