package medium

type TreeInfo struct {
	Seen      int
	LastValue int
}

func NewTreeInfo() *TreeInfo {
	return &TreeInfo{
		Seen:      0,
		LastValue: 0,
	}
}

func FindKthLargestValueInBst(tree *BinaryTreeNode, k int) int {
	treeInfo := NewTreeInfo()
	reverseInOrderTraverse(tree, k, treeInfo)
	return treeInfo.LastValue
}

func reverseInOrderTraverse(tree *BinaryTreeNode, k int, info *TreeInfo) {
	if tree == nil || info.Seen >= k {
		return
	}
	reverseInOrderTraverse(tree.Right, k, info)
	if info.Seen < k {
		info.Seen++
		info.LastValue = tree.Value
		reverseInOrderTraverse(tree.Left, k, info)
	}
}
