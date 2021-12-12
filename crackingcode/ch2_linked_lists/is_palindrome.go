package ch2_linked_lists

func IsPalindrome(head *SinglyLinkedNode) bool {
	reversed := reverseAndClone(head)
	return areEqualLists(head, reversed)
}

func reverseAndClone(node *SinglyLinkedNode) *SinglyLinkedNode {
	var head *SinglyLinkedNode
	for node != nil {
		n := &SinglyLinkedNode{Value: node.Value}
		n.Next = head
		head = n
		node = node.Next
	}
	return head
}

func areEqualLists(first, second *SinglyLinkedNode) bool {
	for first != nil && second != nil {
		if first.Value != second.Value {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return first == nil && second == nil
}
