package easy

import "testing"

func TestIsValidSubsequence(t *testing.T) {
	result := IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10})
	if result == false {
		t.Error("Expected: true")
	}
}
