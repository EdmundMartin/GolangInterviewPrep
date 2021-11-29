package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	palindrome := "tat"
	nonPalindrome := "bobobs"

	if IsPalindrome(palindrome) == false {
		t.Error("Expected: true")
	}
	if IsPalindrome(nonPalindrome) == true {
		t.Error("Expected: false")
	}
}
