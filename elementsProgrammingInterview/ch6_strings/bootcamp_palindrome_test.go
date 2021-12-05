package ch6_strings

import "testing"

func TestIsPalindrome(t *testing.T) {

	if !IsPalindrome("mom") {
		t.Error("Should be palindrome")
	}

	if IsPalindrome("edmund") {
		t.Error("Should not be palindrome")
	}
}