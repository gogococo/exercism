// A package containing 1 function that helps users to share, 
package twofer

// A simple function that takes in a name & outputs
// "One for {name}, one for me."
func ShareWith(name string) string {
	var name_to_print string = "you"
	if name != "" { name_to_print = name}

	return "One for " + name_to_print + ", one for me."
}
