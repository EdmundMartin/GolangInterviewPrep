package easy

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	result := InsertionSort([]int{24, 2, 242, 57, 243, 298, 56, 8, 9})
	if !reflect.DeepEqual(result, []int{2, 8, 9, 24, 56, 57, 242, 243, 298}) {
		t.Errorf("Unexpected ordering got: %v", result)
	}
}
