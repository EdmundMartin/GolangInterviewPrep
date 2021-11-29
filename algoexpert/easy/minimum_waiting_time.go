package easy

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	totalWaitingTime := 0
	for idx, duration := range queries {
		queriesRemaining := len(queries) - (idx + 1)
		totalWaitingTime += queriesRemaining * duration
	}
	return totalWaitingTime
}
