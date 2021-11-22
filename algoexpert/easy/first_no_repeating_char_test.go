package easy

import "testing"

func TestFirstNonRepeatingChar(t *testing.T) {
	testStr := "abcdcaf"
	result := FirstNonRepeatingChar(testStr)
	if result != 1 {
		t.Error("unexpected result")
	}
}