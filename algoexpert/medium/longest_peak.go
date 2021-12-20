package medium


func measurePeak(array []int, idx int) int {
	size := 1
	prev := array[idx]
	backIdx := idx
	frontIdx := idx

	for backIdx > 0 && array[backIdx - 1] < prev {
		size++
		prev = array[backIdx - 1]
		backIdx--
	}
	prev = array[idx]

	for frontIdx < len(array) - 1 && array[frontIdx + 1] < prev {
		size++
		prev = array[frontIdx + 1]
		frontIdx++
	}
	return size
}

func LongestPeak(array []int) int {
	var peaks []int

	for idx, value := range array {
		if idx == 0 || idx == len(array) - 1 {
			continue
		}
		if value > array[idx - 1] && array[idx + 1] < value {
			peaks = append(peaks, idx)
		}
	}
	longest := 0
	for _, p := range peaks {
		longest = MaxInt(longest, measurePeak(array, p))
	}
	return longest
}
