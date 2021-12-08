package ch7_linked_lists

type ListNode struct {
	Data int
	Next *ListNode
}


func SearchInList(l *ListNode, key int) *ListNode {
	for l != nil && l.Data != key {
		l = l.Next
	}
	return l
}

func InsertAfter(node *ListNode, newNode *ListNode)  {
	newNode.Next = node.Next
	node.Next = newNode
}

func DeleteAfter(node *ListNode) {
	node.Next = node.Next.Next
}