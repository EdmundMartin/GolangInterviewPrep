package ch2_linked_lists

type SinglyLinkedNode struct {
	Value int
	Next  *SinglyLinkedNode
}

func LinkedListToSlice(head *SinglyLinkedNode) []int {
	var values []int
	for head != nil {
		values = append(values, head.Value)
		head = head.Next
	}
	return values
}

func LinkedListFromSlice(values []int) *SinglyLinkedNode {
	dummyHead := &SinglyLinkedNode{Value: 0}
	root := dummyHead
	for _, val := range values {
		root.Next = &SinglyLinkedNode{Value: val}
		root = root.Next
	}
	return dummyHead.Next
}
