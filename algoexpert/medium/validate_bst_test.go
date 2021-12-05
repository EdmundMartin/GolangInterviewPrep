package medium

import "testing"

func TestValidateBst(t *testing.T) {
	validTree := &BinaryTreeNode{Value: 10}
	validTree.Left = &BinaryTreeNode{Value: 7}
	validTree.Left.Left = &BinaryTreeNode{Value: 6}
	validTree.Left.Right = &BinaryTreeNode{Value: 7}

	if ValidateBst(validTree) != true {
		t.Error("Tree should be valid")
	}

	invalidTree := &BinaryTreeNode{Value: 10}
	invalidTree.Right = &BinaryTreeNode{Value: 12}
	invalidTree.Right.Left = &BinaryTreeNode{Value: 11}
	invalidTree.Right.Right = &BinaryTreeNode{Value: 8}

	if ValidateBst(invalidTree) == true {
		t.Error("Tree should be invalid")
	}
}
