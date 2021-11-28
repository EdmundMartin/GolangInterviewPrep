package ch1_arrays_strings

import "testing"

func TestPermutationCount(t *testing.T) {
	if PermutationCount("dog", "god") != true {
		t.Error("Expected true, got false")
	}

	if PermutationCount("rita", "iitr") != false {
		t.Error("Expected false, got true")
	}
}

func TestPermutationSort(t *testing.T) {
	if PermutationSort("dog", "god") != true {
		t.Error("Expected true, got false")
	}

	if PermutationSort("rita", "ittr") != false {
		t.Error("Expected false, got true")
	}
}
