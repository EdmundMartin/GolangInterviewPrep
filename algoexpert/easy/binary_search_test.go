package easy

import "testing"

func TestBinarySearch(t *testing.T) {
	testArray := []int{0, 1, 21, 33, 45, 61, 71, 72, 73}
	testTarget := 33
	result := BinarySearch(testArray, testTarget)
	if result != 3 {
		t.Errorf("Expected index 3, got %d", result)
	}
}
