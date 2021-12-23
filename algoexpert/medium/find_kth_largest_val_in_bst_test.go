package medium

import "testing"

func TestFindKthLargestValueInBst(t *testing.T) {
	root := &BinaryTreeNode{Value: 10}
	root.Right = &BinaryTreeNode{Value: 12}
	root.Right.Right = &BinaryTreeNode{Value: 14}
	root.Left = &BinaryTreeNode{Value: 8}
	root.Left.Right = &BinaryTreeNode{Value: 9}

	if FindKthLargestValueInBst(root, 3) != 10 {
		t.Error("Unexpected error")
	}
}
