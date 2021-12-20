package medium

import "testing"

func TestLongestPeak(t *testing.T) {
	result := LongestPeak([]int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3})

	if result != 6 {
		t.Errorf("Unexpected result, got value: %d", result)
	}
}