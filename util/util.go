package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letter = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIODFGHJKXCVBNM")
	ret := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range ret {
		ret[i] = letter[rand.Intn(len(letter))]
	}
	return string(ret)
}
