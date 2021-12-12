package ch6_strings

import "testing"

func TestRomanToInt(t *testing.T) {
	testInput := "XLII"

	if RomanToInt(testInput) != 42 {
		t.Error("Unexpected value")
	}
}
