package medium

import "testing"

func TestFirstDuplicateTwoPointers(t *testing.T) {
	result := FirstDuplicateTwoPointers([]int{2, 1, 5, 2, 3, 3, 4})

	if result != 2 {
		t.Errorf("Got %d", result)
	}
}


func TestFirstDuplicateUsingSet(t *testing.T) {
	result := FirstDuplicateUsingSet([]int{2, 1, 5, 2, 3, 3, 4})

	if result != 2 {
		t.Errorf("Got %d", result)
	}
}