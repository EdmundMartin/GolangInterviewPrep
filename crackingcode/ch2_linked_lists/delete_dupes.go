package ch2_linked_lists

func mapContains(set map[int]interface{}, val int) bool {
	_, ok := set[val]
	return ok
}

func DeleteDuplicates(node *SinglyLinkedNode) *SinglyLinkedNode {
	head := node
	seenValues := make(map[int]interface{})
	var previous *SinglyLinkedNode
	for node != nil {
		if mapContains(seenValues, node.Value) {
			previous.Next = node.Next
		} else {
			seenValues[node.Value] = nil
			previous = node
		}
		node = node.Next
	}
	return head
}

func DeleteDuplicatesNoBuffer(node *SinglyLinkedNode) *SinglyLinkedNode {
	head := node
	current := node
	for current != nil {
		runner := current
		for runner.Next != nil {
			if runner.Next.Value == current.Value {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next
	}
	return head
}
