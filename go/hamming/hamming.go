package hamming

import "errors"

// TODO:
// Convert strings to rune slices first ?

// Distance calculates the Hamming Distance between two Strings.
func Distance(a, b string) (int, error) {
	var HammingCount int = 0

	if len(a) == len(b) {
		if a != b {
			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					HammingCount++
				}
			}
		}
		return HammingCount, nil
	}
	return -0, errors.New("strings are not same length")
}
