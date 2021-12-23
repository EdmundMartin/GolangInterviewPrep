package easy

func SelectionSort(array []int) []int {
	currentIdx := 0
	for currentIdx < len(array)-1 {
		smallestIdx := currentIdx
		for i := currentIdx + 1; i < len(array); i++ {
			if array[smallestIdx] > array[i] {
				smallestIdx = i
			}
		}
		array[currentIdx], array[smallestIdx] = array[smallestIdx], array[currentIdx]
		currentIdx++
	}
	return array
}
