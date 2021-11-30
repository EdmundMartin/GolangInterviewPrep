package easy

import "math"

func ThreeLargestNumbers(array []int) []int  {
	first := math.MinInt32
	second := math.MinInt32
	third := math.MinInt32

	for _, val := range array {
		if val > first {
			first, second, third = val, first, second
		} else if val > second {
			second, third = val, second
		} else if val > third {
			third = val
		}
	}

	return []int{third, second, first}
}