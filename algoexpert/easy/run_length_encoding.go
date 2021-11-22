package easy

import (
	"bytes"
	"fmt"
	"strings"
)

func RunLengthEncoding(target string) string {
	result := bytes.Buffer{}
	prev := strings.Split(target, "")[0]
	count := 1
	for _, current := range strings.Split(target, "")[1:] {
		if current != prev {
			result.WriteString(fmt.Sprintf("%s%d", prev, count))
			count = 0
		}
		if count == 9 {
			result.WriteString(fmt.Sprintf("%s%d", prev, count))
			count = 0
		}
		prev = current
		count += 1
	}
	result.WriteString(fmt.Sprintf("%s%d", prev, count))
	return result.String()
}