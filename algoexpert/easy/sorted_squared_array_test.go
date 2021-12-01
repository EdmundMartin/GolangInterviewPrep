package easy

import (
	"reflect"
	"testing"
)

func TestSortedSquareArray(t *testing.T) {
	array := []int{1, 2, 3, 5, 6, 8, 9}
	solution := SortedSquareArray(array)

	if !reflect.DeepEqual(array,[]int{1, 2, 3, 5, 6, 8, 9}) {
		t.Error("Expected array to be unmodified")
	}

	if !reflect.DeepEqual(solution, []int{1, 4, 9, 25, 36, 64, 81}) {
		t.Error("Got wrong result")
	}
}
