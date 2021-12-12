package ch2_linked_lists

func Partition(node *SinglyLinkedNode, target int) *SinglyLinkedNode {
	head := node
	tail := node

	for node != nil {
		nextNode := node.Next
		if node.Value < target {
			node.Next = head
			head = node
		} else {
			tail.Next = node
			tail = node
		}
		node = nextNode
	}
	tail.Next = nil
	return head
}
