package medium

import (
	"reflect"
	"testing"
)

func TestSearchInSortedMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 12, 15, 1000},
		{2, 5, 19, 31, 32, 1001},
		{3, 8, 24, 33, 35, 1002},
		{40, 41, 42, 44, 45, 1003},
		{99, 100, 103, 106, 128, 1004},
	}
	target := 44

	result := SearchInSortedMatrix(matrix, target)

	if !reflect.DeepEqual(result, []int{3, 3}) {
		t.Errorf("Found error: %v", result)
	}
}
