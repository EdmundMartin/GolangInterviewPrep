package easy

import (
	"bytes"
)

func CaesarCipher(target string, key int32) string {
	buffer := bytes.Buffer{}
	if key > 26 {
		key = key % 26
	}
	for _, letter := range target {
		if (letter + key) > 122 {
			diff := (122 - letter) + 1
			buffer.WriteRune(diff)
		} else {
			buffer.WriteRune(letter + key)
		}
	}
	return buffer.String()
}
