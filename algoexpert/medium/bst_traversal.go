package medium


type BinaryTreeNode struct {
	Value int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}


func InOrder(node *BinaryTreeNode, array *[]int) *[]int {
	if node != nil {
		InOrder(node.Left, array)
		*array = append(*array, node.Value)
		InOrder(node.Right, array)
	}
	return array
}

func PreOrder(node *BinaryTreeNode, array *[]int) *[]int {
	if node != nil {
		*array = append(*array, node.Value)
		PreOrder(node.Left, array)
		PreOrder(node.Right, array)
	}
	return array
}

func PostOrder(node *BinaryTreeNode, array *[]int) *[]int  {
	if node != nil {
		PostOrder(node.Left, array)
		PostOrder(node.Right, array)
		*array = append(*array, node.Value)
	}
	return array
}