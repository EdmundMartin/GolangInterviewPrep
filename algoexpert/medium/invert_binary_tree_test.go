package medium

import (
	"reflect"
	"testing"
)

func getTestTree() *BinaryTreeNode {
	root := &BinaryTreeNode{Value: 10}
	root.Left = &BinaryTreeNode{Value: 7}
	root.Right = &BinaryTreeNode{Value: 12}
	return root
}

func TestInvertBinaryTree(t *testing.T) {
	testTree := getTestTree()
	InvertBinaryTree(testTree)
	var values []int
	result := InOrder(testTree, &values)
	if !reflect.DeepEqual(result, &[]int{12, 10, 7}) {
		t.Errorf("Got result: %v", result)
	}
}

func TestInvertBinaryTreeRecursive(t *testing.T) {
	testTree := getTestTree()
	InvertBinaryTreeRecursive(testTree)
	var values []int
	result := InOrder(testTree, &values)
	if !reflect.DeepEqual(result, &[]int{12, 10, 7}) {
		t.Errorf("Got result: %v", result)
	}
}
