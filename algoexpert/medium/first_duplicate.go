package medium


func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FirstDuplicateTwoPointers(array []int) int {
	minSecondIdx := len(array)

	for idx, value := range array {
		for j := idx + 1; j < len(array); j++ {
			if value == array[j] {
				minSecondIdx = minInt(minSecondIdx, j)
			}
		}
	}

	if minSecondIdx == len(array) {
		return  -1
	}

	return array[minSecondIdx]
}


func FirstDuplicateUsingSet(array []int) int {
	seen := make(map[int]interface{})
	for _, val := range array {
		_, ok := seen[val]
		if ok {
			return val
		}
		seen[val] = nil
	}
	return -1
}