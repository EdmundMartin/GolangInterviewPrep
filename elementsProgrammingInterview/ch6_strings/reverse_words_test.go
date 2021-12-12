package ch6_strings

import (
	"strings"
	"testing"
)

func TestReverseWords(t *testing.T) {
	result := ReverseWords(strings.Split("ram is costly", ""))
	final := strings.Join(result, "")

	if final != "costly is ram" {
		t.Error("Unexpected result")
	}
}
