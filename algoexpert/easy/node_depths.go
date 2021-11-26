package easy


func NodeDepths(root *BinaryNode, depth int) int {
	if root == nil {
		return 0
	}
	return depth + NodeDepths(root.Left, depth + 1) + NodeDepths(root.Right, depth + 1)
}