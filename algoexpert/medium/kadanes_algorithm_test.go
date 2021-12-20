package medium

import "testing"

func TestKadanesAlgorithm(t *testing.T) {
	result := KadanesAlgorithm([]int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4})

	if result != 19 {
		t.Errorf("Unexpected result, got: %d", result)
	}
}