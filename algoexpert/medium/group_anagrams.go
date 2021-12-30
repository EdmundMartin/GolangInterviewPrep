package medium

import (
	"sort"
	"strings"
)

func GroupAnagrams(words []string) [][]string {
	results := make(map[string][]string)
	var output [][]string
	for _, word := range words {
		splitString := strings.Split(word, "")
		sort.Strings(splitString)
		sortedWord := strings.Join(splitString, "")
		val, ok := results[sortedWord]
		if ok {
			val = append(val, word)
			results[sortedWord] = val
		} else {
			results[sortedWord] = []string{word}
		}
	}
	for _, val := range results {
		output = append(output, val)
	}
	return output
}