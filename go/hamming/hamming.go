package hamming

import "errors"

// Distance calculates the Hamming Distance between two Strings.
func Distance(a, b string) (int, error) {
	var HammingCount int = 0
	ar := []rune(a)
	br := []rune(b)

	if len(ar) != len(br) {
		return -0, errors.New("strings are not same length")
	}

	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			HammingCount++
		}
	}
	return HammingCount, nil
}
