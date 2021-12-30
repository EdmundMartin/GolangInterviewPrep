package medium

import (
	"reflect"
	"testing"
)

func TestMinimumCharactersForWords(t *testing.T) {
	inputs := []string{"deed", "dead"}
	result := MinimumCharactersForWords(inputs)
	if !reflect.DeepEqual(result, []string{"a", "d", "d", "e", "e"}) {
		t.Error("Unexpected result")
	}
}
