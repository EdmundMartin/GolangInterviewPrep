package ch1_arrays_strings

import (
	"bytes"
	"fmt"
)

func Compress(target string) string {
	buffer := bytes.Buffer{}
	countRepeating := 0

	for idx, r := range target {
		countRepeating++

		if idx+1 >= len(target) || target[idx] != target[idx+1] {
			buffer.WriteRune(r)
			buffer.WriteString(fmt.Sprintf("%d", countRepeating))
			countRepeating = 0
		}
	}
	compressed := buffer.String()
	if len(compressed) < len(target) {
		return compressed
	}
	return target
}
