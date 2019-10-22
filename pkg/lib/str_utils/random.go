package str_utils

import "math/rand"

const randLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandStr(n int) string {
	b := make([]byte, n)
	lenLetters := len(randLetters)
	for i := range b {
		b[i] = randLetters[rand.Intn(lenLetters)]
	}
	return string(b)
}
