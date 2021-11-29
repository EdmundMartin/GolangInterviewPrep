package easy

func branchSumHelper(node *BinaryNode, runningSum int, sums *[]int) {
	if node == nil {
		return
	}
	newSum := runningSum + node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, newSum)
		return
	}
	branchSumHelper(node.Left, newSum, sums)
	branchSumHelper(node.Right, newSum, sums)
}

func BranchSum(root *BinaryNode) []int {
	var sums []int
	branchSumHelper(root, 0, &sums)
	return sums
}
