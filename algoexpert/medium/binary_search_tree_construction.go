package medium

type BST struct {
	Value int
	Left *BST
	Right *BST
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}


func (b *BST) Insert(value int) *BST {
	currentNode := b
	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = NewBST(value)
				break
			} else {
				currentNode = currentNode.Left
			}
		} else {
			if currentNode.Right == nil {
				currentNode.Right = NewBST(value)
				break
			} else {
				currentNode = currentNode.Right
			}
		}
	}
	return b
}


func (b *BST) Contains(value int) bool {
	currentNode := b
	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return true
		}
	}
	return false
}


func (b *BST) GetMinValue() int {
	currentNode := b
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode.Value
}

func selectLeftIfExists(b *BST) *BST {
	if b.Left != nil {
		return b.Left
	}
	return b.Right
}

func (b *BST) Remove(value int, parentNode *BST) *BST {
	currentNode := b
	for currentNode != nil {
		if value < currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Right
		} else {
			if currentNode.Left != nil && currentNode.Right != nil {
				currentNode.Value = currentNode.Right.GetMinValue()
				currentNode.Right.Remove(currentNode.Value, currentNode)
			} else if parentNode == nil { // Edge case when dealing with root noe
				if currentNode.Left != nil {
					currentNode.Value = currentNode.Left.Value
					currentNode.Right = currentNode.Left.Right
					currentNode.Left = currentNode.Left.Left
				} else if currentNode.Right != nil {
					currentNode.Value = currentNode.Right.Value
					currentNode.Left = currentNode.Right.Left
					currentNode.Right = currentNode.Right.Right
				} else {
					// What to do if removing root node with no children?
					currentNode.Value = 0
				}
			} else if parentNode.Left == currentNode {
				parentNode.Left = selectLeftIfExists(currentNode)
			} else if parentNode.Right == currentNode {
				parentNode.Right = selectLeftIfExists(currentNode)
			}
			break
		}
	}
	return b
}