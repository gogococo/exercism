package hamming

import "errors"

// Distance calculates the Hamming Distance between two Strings.
func Distance(a, b string) (int, error) {
	var HammingCount int = 0

	if len(a) != len(b) {
		return -1, errors.New("Strings are not same length")
	}

	if len(a) == len(b) {
		if a != b {
			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					HammingCount++
				}
			}
		}
	}
	return HammingCount, nil
}
