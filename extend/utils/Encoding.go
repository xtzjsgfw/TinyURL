package utils

import (
	"bytes"
	"math"
)

const dictLength = 62

var dict []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func Base62Encode(id int) string {
	result := make([]byte, 0)
	number := id
	for number > 0 {
		round := number / dictLength
		remain := number % dictLength
		result = append([]byte{dict[remain]}, result...)
		number = round
	}
	return string(result)
}

func Base62Decode(code string) int {
	var result int = 0
	codeLength := len(code)
	for i, c := range []byte(code) {
		result += bytes.IndexByte(dict, c) * int(math.Pow(dictLength, float64(codeLength - 1 - i)))
	}
	return result
}