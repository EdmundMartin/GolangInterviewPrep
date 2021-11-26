package easy

import (
	"reflect"
	"testing"
)

func TestBranchSum(t *testing.T) {
	root := &BinaryNode{Value: 10}
	root.Left = &BinaryNode{Value: 8}
	root.Left.Left = &BinaryNode{Value: 6}
	root.Left.Left.Left = &BinaryNode{Value: 4}
	root.Right = &BinaryNode{Value: 12}
	root.Right.Right = &BinaryNode{Value: 20}

	result := BranchSum(root)
	if !reflect.DeepEqual([]int{28, 42}, result) {
		t.Errorf("Did not get expected result")
	}
}
