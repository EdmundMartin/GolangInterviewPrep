package easy

type BinaryNode struct {
	Value int
	Left *BinaryNode
	Right *BinaryNode
}

func NewBinaryNode(value int) *BinaryNode {
	return &BinaryNode{Value: value}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosest(target, old, new int) int {
	if abs(target - old) < abs(target - new) {
		return old
	}
	return new
}


func FindClosestValue(root *BinaryNode, target int) int {
	closest := root.Value
	node := root
	for node != nil {
		if node.Value == target {
			return target
		} else if target < node.Value {
			closest = findClosest(target, closest, node.Value)
			node = node.Left
		} else {
			closest = findClosest(target, closest, node.Value)
			node = node.Right
		}
	}
	return closest
}