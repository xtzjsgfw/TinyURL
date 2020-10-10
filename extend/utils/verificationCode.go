package utils

import (
	"math/rand"
	"time"
)

var randomCodes = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func CreateVerificationCode() string {
	var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	var pwd []byte = make([]byte, 6)

	for j := 0; j < 6; j++ {
		index := r.Int() % len(randomCodes)

		pwd[j] = randomCodes[index]
	}
	return string(pwd)
}