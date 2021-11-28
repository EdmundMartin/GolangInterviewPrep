package ch2_linked_lists

/*
Implement algo to delete a node in the middle of a singly linked list,
given only access to that node.

Say if you can delete the node or not.
*/

func DeleteNode(node *SinglyLinkedNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	next := node.Next
	node.Value = next.Value
	node.Next = next.Next
	return true
}
