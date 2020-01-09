package reverse

func Reverse(str string) string {
	n := len(str)
	runes := make([]rune, n)

	for _, rune := range str {
		n--
		runes[n] = rune
	}

	return string(runes[n:])
}
