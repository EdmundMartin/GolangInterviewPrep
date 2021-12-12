package easy

import "testing"

func TestClassPhoto(t *testing.T) {
	reds := []int{5, 8, 1, 3, 4}
	blues := []int{6, 9, 2, 4, 5}

	if !ClassPhoto(reds, blues) {
		t.Error("Expected true")
	}
}
