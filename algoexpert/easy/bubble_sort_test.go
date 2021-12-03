package easy

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testArray := []int{11, 2, 4, 8, 9, 3, 1, 1, 139, 273, 15}
	result := BubbleSort(testArray)

	if !reflect.DeepEqual(result, []int{1, 1, 2, 3, 4, 8, 9, 11, 15, 139, 273}) {
		t.Error("Unexpected result")
	}
}
