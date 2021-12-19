package medium


func InvertBinaryTree(tree *BinaryTreeNode) {
	queue := []*BinaryTreeNode{tree}
	for len(queue) > 0 {
		var x *BinaryTreeNode
		x, queue = queue[len(queue) - 1], queue[:len(queue) - 1]
		if x == nil {
			continue
		}
		x.Left, x.Right = x.Right, x.Left
		queue = append(queue, x.Left)
		queue = append(queue, x.Right)
	}
}


func InvertBinaryTreeRecursive(tree *BinaryTreeNode) {
	if tree == nil {
		return
	}
	tree.Left, tree.Right = tree.Right, tree.Left
	InvertBinaryTreeRecursive(tree.Left)
	InvertBinaryTreeRecursive(tree.Right)
}