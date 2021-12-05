package ch2_linked_lists

type TailResult struct {
	Tail *SinglyLinkedNode
	Size int
}

func FindIntersection(first, second *SinglyLinkedNode) *SinglyLinkedNode {
	if first == nil || second == nil {
		return nil
	}
	firstResult := getTailAndSize(first)
	secondResult := getTailAndSize(second)

	if firstResult.Tail != secondResult.Tail {
		return nil
	}

	var shorter *SinglyLinkedNode
	var longer *SinglyLinkedNode
	if firstResult.Size < secondResult.Size {
		shorter, longer = first, second
	} else {
		longer, shorter = first, second
	}

	longer = getKthNode(longer)

	for shorter != longer {
		shorter = shorter.Next
		longer = longer.Next
	}

	return longer
}

func getTailAndSize(node *SinglyLinkedNode) TailResult {
	if node == nil {
		return TailResult{
			Tail: nil,
			Size: 0,
		}
	}
	size := 1
	current := node
	for current.Next != nil {
		size++
		current = current.Next
	}
	return TailResult{
		Tail: current,
		Size: size,
	}
}


func getKthNode(node *SinglyLinkedNode, k int) *SinglyLinkedNode {
	current := node
	for k > 0 && current != nil {
		current = current.Next
		k--
	}
	return current
}