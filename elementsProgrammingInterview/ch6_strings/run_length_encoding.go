package ch6_strings

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)


func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func multChar(s string, count int) string {
	buffer := bytes.Buffer{}
	for i := 0; i < count; i++ {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func EncodingRunLength(s string) string {
	buffer := bytes.Buffer{}
	count := 1
	asSlice := strings.Split(s, "")

	for i := 1; i < len(s); i++ {
		if i == len(s) || asSlice[i] != asSlice[i - 1] {
			buffer.WriteString(fmt.Sprintf("%d%s", count, asSlice[i - 1]))
			count = 1
		} else {
			count++
		}
	}
	return buffer.String()
}


func DecodingRunLength(s string) string {
	count := 0
	result := bytes.Buffer{}

	for _, ch := range strings.Split(s, "") {
		if isInt(ch) {
			val, _ := strconv.Atoi(ch)
			count = (count * 10) + val
		} else {
			result.WriteString(multChar(s, count))
			count = 0
		}
	}
	return result.String()
}