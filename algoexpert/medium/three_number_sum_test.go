package medium

import (
	"reflect"
	"testing"
)

func TestThreeNumberSum(t *testing.T) {
	results := ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0)
	expected := [][]int{
		{-8, 2, 6},
		{-8, 3, 5},
		{-6, 1, 5},
	}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Unexpected value: %v", results)
	}
}
