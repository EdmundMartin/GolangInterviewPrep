package ch1_arrays_strings

import "testing"

func TestCompress(t *testing.T) {
	testString := "eeeddddaaaabcd"
	result := Compress(testString)

	if result != "e3d4a4b1c1d1" {
		t.Errorf("Got unexpected result: %s", result)
	}
}
