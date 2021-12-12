package medium

import (
	"reflect"
	"testing"
)

func TestReverseRange(t *testing.T) {
	test := []string{"h", "e", "l", "l", "o"}
	ReverseRange(&test, 0, len(test)-1)
	if !reflect.DeepEqual(test, []string{"o", "l", "l", "e", "h"}) {
		t.Error("Reverse did not work")
	}
}

func TestReverseWordsInString(t *testing.T) {
	result := ReverseWordsInString("edmund is the best")

	if result != "best the is edmund" {
		t.Errorf("Got unexpected result: %s", result)
	}
}
