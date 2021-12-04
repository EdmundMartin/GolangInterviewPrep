package medium

type LinkedList struct {
	Value int
	Next *LinkedList
}

func SliceToLinkedList(values []int) *LinkedList {
	dummyHead := &LinkedList{Value: 0}
	root := dummyHead
	for _, val := range values {
		root.Next = &LinkedList{Value: val}
		root = root.Next
	}
	return dummyHead.Next
}

func LinkedListToSlice(head *LinkedList) []int {
	var values []int
	for head != nil {
		values = append(values, head.Value)
		head = head.Next
	}
	return values
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	counter := 0
	slow := head
	fast := head
	for counter < k {
		fast = fast.Next
		counter++
	}
	if fast == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
}