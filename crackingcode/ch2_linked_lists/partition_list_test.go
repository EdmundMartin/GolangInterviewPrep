package ch2_linked_lists

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	testData := LinkedListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})

	result := Partition(testData, 7)
	asSlice := LinkedListToSlice(result)
	if !reflect.DeepEqual(asSlice, []int{6, 5, 4, 3, 2, 1, 7, 8, 9, 10, 11}) {
		t.Errorf("Did not get expected value, got %v", asSlice)
	}
}