package ch6_strings

import "testing"

func TestIsPalindromePunctuation(t *testing.T) {
	isPalindrome := "Able was I, ere I saw Elba"
	isNotPalindrome := "Ray a Ray"

	if !IsPalindromePunctuation(isPalindrome) {
		t.Error("Expected palindrome")
	}

	if IsPalindromePunctuation(isNotPalindrome) {
		t.Error("Did not expected palindrome")
	}
}
