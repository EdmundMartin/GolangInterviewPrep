package easy

import "strings"

func FirstNonRepeatingChar(target string) int {
	counter := make(map[string]int)
	charSlice := strings.Split(target, "")
	for _, char := range charSlice {
		_, ok := counter[char]
		if ok {
			counter[char]++
		} else {
			counter[char] = 1
		}
	}
	for idx, char := range charSlice {
		if counter[char] == 1 {
			return idx
		}
	}
	return -1
}