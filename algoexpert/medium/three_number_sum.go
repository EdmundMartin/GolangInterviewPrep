package medium

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	var results [][]int

	for idx := 0; idx < len(array) - 2; idx++ {
		left := idx + 1
		right := len(array) - 1
		for left < right {
			currentSum := array[idx] + array[left] + array[right]
			if currentSum == target {
				results = append(results, []int{
					array[idx],
					array[left],
					array[right],
				})
				left++
				right--
			} else if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return results
}