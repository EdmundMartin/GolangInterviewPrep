package easy

import "testing"

func TestNonConstructibleChange(t *testing.T) {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	result := NonConstructibleChange(coins)

	if result != 20 {
		t.Errorf("Expected 20 got: %d", result)
	}
}