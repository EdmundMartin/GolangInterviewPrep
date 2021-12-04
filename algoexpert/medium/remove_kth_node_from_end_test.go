package medium

import (
	"reflect"
	"testing"
)

func TestRemoveKthNodeFromEnd(t *testing.T) {
	testArray := SliceToLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	RemoveKthNodeFromEnd(testArray, 3)
	array := LinkedListToSlice(testArray)
	if !reflect.DeepEqual(array, []int{1, 2, 3, 4, 5, 6, 7, 9, 10}) {
		t.Error("Unexpected result")
	}
}