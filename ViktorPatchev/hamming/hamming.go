package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	//Distance in DNA
	var counter int
	var err error
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				counter++
			}
		}
	} else {
		err = errors.New("Strings are of different size")
	}
	return counter, err
}
