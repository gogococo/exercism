package raindrops

import "strconv"

// Convert will check the given int for factors 3, 5 or 7 and return "Pling", "Plang" or "Plong" respectively
func Convert(n int) string {
	finalString := ""

	if n%3 == 0 {
		finalString += "Pling"
	}
	if n%5 == 0 {
		finalString += "Plang"
	}
	if n%7 == 0 {
		finalString += "Plong"
	}
	if finalString == "" {
		finalString = strconv.Itoa(n)
	}
	return finalString
}
