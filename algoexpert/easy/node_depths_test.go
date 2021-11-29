package easy

import "testing"

func TestNodeDepths(t *testing.T) {
	root := &BinaryNode{Value: 10}
	root.Left = &BinaryNode{Value: 9}
	root.Right = &BinaryNode{Value: 11}
	root.Left.Left = &BinaryNode{Value: 8}
	root.Right.Right = &BinaryNode{Value: 12}
	root.Right.Right.Right = &BinaryNode{Value: 14}

	result := NodeDepths(root, 0)
	if result != 9 {
		t.Errorf("Expected 9, got %d", result)
	}
}
