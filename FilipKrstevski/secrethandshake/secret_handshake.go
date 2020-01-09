package secret

var steps = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(n uint) []string {
	var handshake = make([]string, 0)

	if n <= 0 {
		return handshake

	}

	for s, step := range steps {
		if (1<<uint(s))&n > 0 {
			handshake = append(handshake, step)
		}
	}
	if (1<<uint(len(steps)))&n > 0 {
		reverse(handshake)

	}
	return handshake
}

func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}

}
