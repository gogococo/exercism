package raindrops

import "strconv"

// Convert will check the given int for factors 3, 5 or 7 and return "Pling", "Plang" or "Plong" respectively
func Convert(n int) string {
	FinalString := ""

	if n%3 == 0 {
		FinalString += "Pling"
	}
	if n%5 == 0 {
		FinalString += "Plang"
	}
	if n%7 == 0 {
		FinalString += "Plong"
	}
	if FinalString == "" {
		FinalString = strconv.Itoa(n)
	}
	return FinalString
}
