package ch1_arrays_strings

import "testing"

func TestOneEditAway(t *testing.T) {
	if !OneEditAway("pale", "ple") {
		t.Error("Expected true")
	}
	if !OneEditAway("pales", "pale") {
		t.Error("Expected true")
	}
	if !OneEditAway("pale", "bale") {
		t.Error("Expected true")
	}
	if OneEditAway("pale", "bae") {
		t.Error("Expected false")
	}
}