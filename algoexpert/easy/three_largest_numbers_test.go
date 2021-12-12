package easy

import (
	"reflect"
	"testing"
)

func TestThreeLargestNumbers(t *testing.T) {
	result := ThreeLargestNumbers([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7})

	if !reflect.DeepEqual(result, []int{18, 141, 541}) {
		t.Error("Got unexpected result")
	}
}
