package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = fmt.Sprintf("%s%d", RandStringBytes(2), rand.Intn(999))
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
