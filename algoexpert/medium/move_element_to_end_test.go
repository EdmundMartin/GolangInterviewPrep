package medium

import (
	"reflect"
	"testing"
)

func TestMoveElementToEnd(t *testing.T) {
	testInput := []int{2, 1, 2, 2, 2, 3, 4, 2}
	target := 2

	result := MoveElementToEnd(testInput, target)
	if !reflect.DeepEqual(result, []int{4, 1, 3, 2, 2, 2, 2, 2}) {
		t.Error("Unexpected value")
	}
}