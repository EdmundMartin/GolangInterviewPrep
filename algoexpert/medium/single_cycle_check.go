package medium

func getNextIdx(array []int, currentIdx int) int {
	jumpSize := array[currentIdx]
	nextIdx := (currentIdx + jumpSize) % len(array)
	if nextIdx >= 0 {
		return nextIdx
	} else {
		return nextIdx + len(array)
	}
}

func HasSingleCycle(array []int) bool {
	elementsVisited := 0
	idx := 0
	for elementsVisited < len(array) {
		if elementsVisited > 0 && idx == 0 {
			return false
		}
		elementsVisited++
		idx = getNextIdx(array, idx)
	}
	return idx == 0
}
