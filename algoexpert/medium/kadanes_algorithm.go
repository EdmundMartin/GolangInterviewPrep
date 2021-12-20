package medium

import "math"


func KadanesAlgorithm(array []int) int {
	currentSums := make([]int, len(array))
	max := math.MinInt32
	for idx, number := range array {
		if idx == 0 {
			currentSums[0] = number
			max = MaxInt(number, max)
		} else {
			newSum := number + currentSums[idx - 1]
			runningMax := MaxInt(number, newSum)
			max = MaxInt(max, runningMax)
			currentSums[idx] = runningMax
		}
	}
	return max
}