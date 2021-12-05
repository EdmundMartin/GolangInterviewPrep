package medium

import "math"

func validationHelper(tree *BinaryTreeNode, minValue int, maxValue int) bool {
	if tree == nil {
		return true
	}
	if tree.Value < minValue || tree.Value >= maxValue {
		return false
	}
	leftValid := validationHelper(tree.Left, minValue, tree.Value)
	rightValid := validationHelper(tree.Right, tree.Value, maxValue)
	return leftValid && rightValid
}

func ValidateBst(tree *BinaryTreeNode) bool {
	return validationHelper(tree, math.MinInt32, math.MaxInt32)
}
