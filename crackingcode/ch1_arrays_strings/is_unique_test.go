package ch1_arrays_strings

import (
	"fmt"
	"testing"
)

func TestIsUnique(t *testing.T) {
	uniqueString := "aedfghjkl"
	duplicateString := "abcdeabd"

	if IsUnique(uniqueString) != true {
		t.Error("Expected true, got false")
	}
	if IsUnique(duplicateString) != false {
		t.Error("Expected false, got true")
	}
}

func TestBlah(t *testing.T) {
	value := "h o"
	for _, r := range value {
		fmt.Println(int(r))
	}
	return
}
