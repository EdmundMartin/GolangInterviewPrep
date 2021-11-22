package easy


func helper(array []int, target, left, right int) int {
	if left > right {
		return - 1
	}
	middle := (left + right) / 2
	if target == array[middle] {
		return middle
	} else if target > array[middle] {
		return helper(array, target, middle + 1, right)
	} else {
		return helper(array, target, left, middle - 1)
	}
}


func BinarySearch(array []int, target int) int {
	return helper(array, target, 0, len(array) - 1)
}

