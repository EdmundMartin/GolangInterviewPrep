package ch2_linked_lists

import (
	"testing"
)

func TestNthToLast(t *testing.T) {
	testData := LinkedListFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result := NthToLast(testData, 3)

	if result.Value != 8 {
		t.Error("Unexpected value")
	}
}
