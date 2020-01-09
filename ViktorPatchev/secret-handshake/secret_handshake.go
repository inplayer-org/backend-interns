package secret

import (
	"strconv"
	"strings"
)

func Handshake(n uint) []string {
	handshake := []string{" ", "wink", "double blink", "close your eyes", "jump"}
	binary := strconv.FormatInt(int64(n), 2)
	final := []string{}
	bin := strings.Split(string(binary), "")
	pom := 0
	for i := len(bin) - 1; i >= 0; i-- {
		pom++
		j, _ := strconv.Atoi(bin[i])
		if j == 1 {
			if pom > 4 {
				reverse(final)
				break
			}
			final = append(final, handshake[pom])
		}
	}
	return final
}

func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}
