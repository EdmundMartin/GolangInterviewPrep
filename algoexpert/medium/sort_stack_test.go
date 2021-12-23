package medium

import (
	"reflect"
	"testing"
)

func TestSortStack(t *testing.T) {
	stack := NewStack(-4, 3, 5, 7, 2, 1, 8)
	SortStack(stack)
	if !reflect.DeepEqual(stack.Stack, []int{-4, 1, 2, 3, 5, 7, 8}) {
		t.Errorf("Got unexpected value: %v", stack.Stack)
	}
}
