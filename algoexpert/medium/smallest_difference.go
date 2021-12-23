package medium

import (
	"math"
	"sort"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SmallestDifference(first, second []int) []int {
	sort.Ints(first)
	sort.Ints(second)
	firstIdx, secondIdx := 0, 0
	smallestDiff := math.MaxInt32
	var currentPair []int

	for firstIdx < len(first) && secondIdx < len(second) {
		difference := Abs(first[firstIdx] - second[secondIdx])
		if difference < smallestDiff {
			currentPair = []int{first[firstIdx], second[secondIdx]}
			smallestDiff = difference
		}
		if first[firstIdx] < second[secondIdx] {
			firstIdx++
		} else if first[firstIdx] > second[secondIdx] {
			secondIdx++
		} else {
			return currentPair
		}
	}
	return currentPair
}
