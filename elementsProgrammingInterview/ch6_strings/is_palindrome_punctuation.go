package ch6_strings

import "strings"

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func IsPalindromePunctuation(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for !isLetter(rune(s[i])) && i < j {
			i++
		}
		for !isLetter(rune(s[j])) && i < j {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
