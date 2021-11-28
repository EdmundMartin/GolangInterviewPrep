package ch2_linked_lists

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	testList := LinkedListFromSlice([]int{1, 1, 1, 3, 3, 4, 4, 5, 6, 7, 7, 8, 8})
	result := DeleteDuplicates(testList)

	listOne := LinkedListToSlice(result)
	if !reflect.DeepEqual(listOne, []int{1, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("Expected lists do not align: got %v", listOne)
	}
}

func TestDeleteDuplicatesNoBuffer(t *testing.T) {
	testList := LinkedListFromSlice([]int{1, 1, 1, 3, 3, 4, 4, 5, 6, 7, 7, 8, 8})
	result := DeleteDuplicatesNoBuffer(testList)

	listOne := LinkedListToSlice(result)
	if !reflect.DeepEqual(listOne, []int{1, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("Expected lists do not align: got %v", listOne)
	}
}
