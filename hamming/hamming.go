package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("")
	}

	var count int
	for idx := range a {
		if a[idx] != b[idx] {
			count++
		}
	}

	return count, nil
}

