package ch2_linked_lists

import "fmt"

type SinglyLinkedNode struct {
	Value int
	Next  *SinglyLinkedNode
}

func (s *SinglyLinkedNode) String() string {
	var vals []int
	node := s
	for node != nil {
		vals = append(vals, node.Value)
		node = node.Next
	}
	return fmt.Sprintf("%v", vals)
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
