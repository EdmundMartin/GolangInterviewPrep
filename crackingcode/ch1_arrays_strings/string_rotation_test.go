package ch1_arrays_strings

import "testing"

func TestIsRotation(t *testing.T) {
	original := "waterbottle"
	rotation := "erbottlewat"

	result := IsRotation(original, rotation)

	if result != true {
		t.Error("Expected valid rotation")
	}
}
