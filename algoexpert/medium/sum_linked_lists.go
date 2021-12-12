package medium


func nodesValue(list *LinkedList) int {
	if list == nil {
		return 0
	}
	return list.Value
}

func nodesNext(list *LinkedList) *LinkedList {
	if list == nil {
		return nil
	}
	return list.Next
}

func SumLinkedLists(firstList, secondList *LinkedList) *LinkedList {
	dummy := &LinkedList{Value: 0}
	currentNode := dummy
	carry := 0
	first, second := firstList, secondList
	for first != nil || second != nil || carry != 0 {
		firstVal := nodesValue(first)
		secondVal := nodesValue(second)
		sumValues := firstVal + secondVal + carry
		newValue := sumValues % 10
		newNode := &LinkedList{Value: newValue}
		currentNode.Next = newNode
		currentNode = newNode
		carry = sumValues / 10
		first = nodesNext(first)
		second = nodesNext(second)
	}
	return dummy.Next
}