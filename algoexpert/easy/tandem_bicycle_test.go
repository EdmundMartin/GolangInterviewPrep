package easy

import "testing"

func TestTandemBicycle(t *testing.T) {
	red := []int{5, 5, 3, 9, 2}
	blue := []int{3, 6, 7, 2, 1}

	result := TandemBicycle(red, blue, true)

	if result != 32 {
		t.Errorf("Expected 32, got: %d", result)
	}
}
