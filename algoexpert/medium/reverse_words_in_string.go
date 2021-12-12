package medium

import "strings"

func ReverseRange(contents *[]string, start, end int) {
	for start < end {
		(*contents)[start], (*contents)[end] = (*contents)[end], (*contents)[start]
		start++
		end--
	}
}

func ReverseWordsInString(value string) string {
	characters := strings.Split(value, "")
	ReverseRange(&characters, 0, len(characters)-1)

	startIdx := 0
	for startIdx < len(characters) {
		endIdx := startIdx
		for endIdx < len(characters) && characters[endIdx] != " " {
			endIdx++
		}
		ReverseRange(&characters, startIdx, endIdx-1)
		startIdx = endIdx + 1
	}

	return strings.Join(characters, "")
}
