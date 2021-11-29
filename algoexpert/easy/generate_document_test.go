package easy

import "testing"

func TestGenerateDocument(t *testing.T) {
	testDocument := "Edmund Martin"
	testCharsTrue := "dmundE artinM"
	testCharsFalse := "dmunde artin"

	if GenerateDocument(testCharsTrue, testDocument) == false {
		t.Error("Expected: true")
	}

	if GenerateDocument(testCharsFalse, testDocument) == true {
		t.Error("Expected: false")
	}
}