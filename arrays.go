package generator

import (
	"math/rand"
	"time"
)

func shuffle(src []string) []string {
	dest := src
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := len(dest) - 1; i > 0; i-- {
		j := rand.Intn(i)
		dest[i], dest[j] = dest[j], dest[i]
	}
	return dest
}

func random(src []string) string {
	return src[rand.Intn(len(src))]
}
