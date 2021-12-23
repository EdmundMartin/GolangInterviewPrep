package medium

import (
	"reflect"
	"testing"
)

func TestSumLinkedLists(t *testing.T) {
	first := SliceToLinkedList([]int{1, 3, 4})
	second := SliceToLinkedList([]int{4, 4, 5})

	output := SumLinkedLists(first, second)
	result := LinkedListToSlice(output)
	if !reflect.DeepEqual(result, []int{5, 7, 9}) {
		t.Errorf("Unexpected result, got: %v", result)
	}
}
