package easy

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testArray := []int{3, 5, -4, 8, 11, 1, -1, 6}
	result := TwoSum(testArray, 10)
	if !reflect.DeepEqual(result, []int{11, -1}) {
		t.Errorf("Got an unexpected error: %v", result)
	}
}