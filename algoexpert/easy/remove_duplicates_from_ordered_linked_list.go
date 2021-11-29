package easy

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(head *LinkedList) *LinkedList {
	var prev *LinkedList = nil
	currentNode := head
	for currentNode != nil {
		if prev != nil && prev.Value != currentNode.Value {
			prev.Next = currentNode
			prev = currentNode
		}
		if prev != nil && currentNode.Next == nil {
			if currentNode.Value == prev.Value {
				prev.Next = nil
			}
		}
		if prev == nil {
			prev = currentNode
		}
		currentNode = currentNode.Next
	}
	return head
}

func RemoveDuplicatesFromLinkedListAlt(head *LinkedList) *LinkedList {
	currentNode := head
	for currentNode != nil {
		nextNode := currentNode.Next
		for nextNode != nil && nextNode.Value == currentNode.Value {
			nextNode = nextNode.Next
		}

		currentNode.Next = nextNode
		currentNode = nextNode
	}
	return head
}
