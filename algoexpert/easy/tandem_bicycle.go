package easy

import (
	"sort"
)

func TandemBicycle(redSpeeds, blueSpeeds []int, fastest bool) int {
	sort.Ints(redSpeeds)
	sort.Ints(blueSpeeds)

	if fastest == false {
		redSpeeds = reverseInPlace(redSpeeds)
	}

	total := 0
	for idx, _ := range redSpeeds {
		first := redSpeeds[idx]
		second := blueSpeeds[len(blueSpeeds)-idx-1]
		total += maxInt(first, second)
	}

	return total
}

func reverseInPlace(target []int) []int {
	start := 0
	end := len(target) - 1
	for start < end {
		target[start], target[end] = target[end], target[start]
		start++
		end--
	}
	return target
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
