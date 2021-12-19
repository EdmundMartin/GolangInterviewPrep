package medium

import (
	"reflect"
	"testing"
)

func TestSmallestDifference(t *testing.T) {
	result := SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})

	if !reflect.DeepEqual(result, []int{28, 26}) {
		t.Errorf("Unexpected result: %v", result)
	}
}
