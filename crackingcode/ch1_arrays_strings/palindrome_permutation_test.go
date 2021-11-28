package ch1_arrays_strings

import "testing"

func TestPalindromePermutationHashTable(t *testing.T) {
	if PalindromePermutationHashTable("tact coa") != true {
		t.Error("Expected true, got false")
	}

	if PalindromePermutationHashTable("edmund martin") != false {
		t.Error("Expected false, got true")
	}
}

func TestPalindromePermutationOptimised(t *testing.T) {
	if PalindromePermutationOptimised("tact coa") != true {
		t.Error("Expected true, got false")
	}

	if PalindromePermutationOptimised("edmund martin") != false {
		t.Error("Expected false, got true")
	}
}
