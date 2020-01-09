package isogram

import "strings"

func IsIsogram(in string) (res bool) {
	var count map[rune]int = make(map[rune]int)
	w := strings.Replace(strings.ToLower(in), "-", "", -1)
	w = strings.Replace(strings.ToLower(w), " ", "", -1)
	for _, i := range w {
		count[i]++
	}
	for _, i := range count {
		if i > 1 {
			return false
		}

	}
	return true
}
