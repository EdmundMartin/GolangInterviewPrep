package easy

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	res := SelectionSort([]int{10, 27, 7, 8, 3, 5, 4, 6})

	if !reflect.DeepEqual(res, []int{3, 4, 5, 6, 7, 8, 10, 27}) {
		t.Errorf("Unexpected result, got: %v", res)
	}
}
