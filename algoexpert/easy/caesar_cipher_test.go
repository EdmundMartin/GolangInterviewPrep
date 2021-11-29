package easy

import (
	"testing"
)

func TestCaesarCipher(t *testing.T) {
	testData := "edmund"
	result := CaesarCipher(testData, 3)
	if result != "hgpxqg" {
		t.Error("Unexpected result")
	}
}
