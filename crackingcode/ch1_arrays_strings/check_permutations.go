package ch1_arrays_strings

import (
	"sort"
	"strings"
)

func sortString(target string) string {
	chars := strings.Split(target, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}


func PermutationSort(first, second string) bool  {
	if len(first) != len(second) {
		return false
	}
	return sortString(first) == sortString(second)
}

func PermutationCount(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	charCounts := [128]int{}

	for _, charRune := range first {
		charCounts[int(charRune)]++
	}
	for _, charRune := range second {
		charCounts[int(charRune)]--
		if charCounts[int(charRune)] < 0 {
			return false
		}
	}
	return true
}