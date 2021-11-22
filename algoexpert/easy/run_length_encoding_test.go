package easy

import (
	"testing"
)

func TestRunLengthEncoding(t *testing.T) {
	testTarget := "AAAAAAAAAAAAAAAAAAAAAHHHHCCCCDDDDPPPSL"
	result := RunLengthEncoding(testTarget)
	if result != "A9A9A3H4C4D4P3S1L1" {
		t.Error("Did not get expected result")
	}
}
