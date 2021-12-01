package easy

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	changeCreated := 0
	for _, coin := range coins {
		if coin > changeCreated + 1 {
			return changeCreated + 1
		}
		changeCreated += coin
	}
	return changeCreated + 1
}