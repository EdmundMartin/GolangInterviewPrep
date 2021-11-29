package easy

import (
	"reflect"
	"testing"
)

func SliceToLinkedList(values []int) *LinkedList {
	dummyHead := &LinkedList{Value: 0}
	root := dummyHead
	for _, val := range values {
		root.Next = &LinkedList{Value: val}
		root = root.Next
	}
	return dummyHead.Next
}

func LinkedListToSlice(head *LinkedList) []int {
	var values []int
	for head != nil {
		values = append(values, head.Value)
		head = head.Next
	}
	return values
}

func TestRemoveDuplicatesFromLinkedList(t *testing.T) {
	testList := SliceToLinkedList([]int{1, 2, 3, 4, 4, 5, 5, 6, 6, 8, 9})
	expectedList := []int{1, 2, 3, 4, 5, 6, 8, 9}
	result := RemoveDuplicatesFromLinkedList(testList)

	if !reflect.DeepEqual(expectedList, LinkedListToSlice(result)) {
		t.Errorf("Expected matching results, got: %v and %v", expectedList, result)
	}
}

func TestRemoveDuplicatesFromLinkedListAlt(t *testing.T) {
	testList := SliceToLinkedList([]int{1, 2, 3, 4, 4, 5, 5, 6, 6, 8, 9})
	expectedList := []int{1, 2, 3, 4, 5, 6, 8, 9}
	result := RemoveDuplicatesFromLinkedListAlt(testList)

	if !reflect.DeepEqual(expectedList, LinkedListToSlice(result)) {
		t.Errorf("Expected matching results, got: %v and %v", expectedList, result)
	}
}
