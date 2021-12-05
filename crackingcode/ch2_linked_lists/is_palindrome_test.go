package ch2_linked_lists

import "testing"

func TestIsPalindrome(t *testing.T) {
	isPalindrome := LinkedListFromSlice([]int{3, 7, 3})
	notPalindrome := LinkedListFromSlice([]int{3, 7, 3, 8})
	if IsPalindrome(isPalindrome) == false {
		t.Error("Expected true")
	}

	if IsPalindrome(notPalindrome) == true {
		t.Error("Expected false")
	}
}