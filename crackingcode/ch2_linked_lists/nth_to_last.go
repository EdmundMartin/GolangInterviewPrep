package ch2_linked_lists

func NthToLast(head *SinglyLinkedNode, k int) *SinglyLinkedNode {
	fast := head
	slow := head

	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
