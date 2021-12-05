package ch6_strings


func IsPalindrome(str string) bool {
	for i := 0; i < len(str) / 2; i++ {
		if str[i] != str[len(str) - 1 - i] {
			return false
		}
	}
	return true
}

