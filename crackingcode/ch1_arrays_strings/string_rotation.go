package ch1_arrays_strings

import "strings"

func isSubstring(first, second string) bool {
	return strings.Contains(first, second)
}

func IsRotation(first, second string) bool {
	length := len(first)

	if length == len(second) && length > 0 {
		firstRepeated := first + first
		return isSubstring(firstRepeated, second)
	}
	return false
}
