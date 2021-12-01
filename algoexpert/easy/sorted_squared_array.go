package easy

import "sort"

func SortedSquareArray(array []int) []int {
	result := make([]int, len(array))

	for idx, value := range array {
		newValue := value * value
		result[idx] = newValue
	}
	sort.Ints(result)
	return result
}