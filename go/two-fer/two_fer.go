// Package twofer implements a function for sharing.
package twofer

// ShareWith helps you to share another person.
func ShareWith(name string) string {
	var nameToPrint string = "you"
	if name != "" {
		nameToPrint = name
	}

	return "One for " + nameToPrint + ", one for me."
}
