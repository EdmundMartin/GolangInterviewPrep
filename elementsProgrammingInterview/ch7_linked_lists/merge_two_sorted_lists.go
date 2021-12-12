package ch7_linked_lists

func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{
		Data: 0,
		Next: nil,
	}
	tail := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Data <= l2.Data {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummyHead.Next
}
