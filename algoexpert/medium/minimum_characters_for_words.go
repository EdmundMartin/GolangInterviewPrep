package medium

import "strings"

func countLettersInWord(word string) map[string]int {
	wordSlice := strings.Split(word, "")
	result := make(map[string]int)
	for _, char := range wordSlice {
		_, ok := result[char]
		if ok {
			result[char]++
		} else {
			result[char] = 1
		}
	}
	return result
}

func MinimumCharactersForWords(words []string) []string {
	globalCount := make(map[string]int)

	for _, word := range words {
		result := countLettersInWord(word)
		for key, value := range result {
			currentVal, ok := globalCount[key]
			if !ok {
				globalCount[key] = value
			} else {
				globalCount[key] = MaxInt(value, currentVal)
			}
		}
	}
	var output []string
	for key, value := range globalCount {
		for i := 0; i < value; i++ {
			output = append(output, key)
		}
	}
	return output
}
