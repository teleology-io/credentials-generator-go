package generator

import (
	"math/rand"
	"time"
)

func shuffle(src []string) []string {
	dest := make([]string, len(src))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func random(src []string) string {
	return src[rand.Intn(len(src))]
}
