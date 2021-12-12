package ch6_strings

func reverseRange(s []string, start, finish int) []string {
	for start < finish {
		s[start], s[finish] = s[finish], s[start]
		start = start + 1
		finish = finish - 1
	}
	return s
}

func ReverseWords(s []string) []string {
	s = reverseRange(s, 0, len(s)-1)

	start := 0
	for {
		finish := start
		for finish < len(s) && s[finish] != " " {
			finish++
		}
		if finish == len(s) {
			break
		}
		s = reverseRange(s, start, finish-1)
		start = finish + 1
	}

	return reverseRange(s, start, len(s)-1)
}
