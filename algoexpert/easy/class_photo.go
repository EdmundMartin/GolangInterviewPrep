package easy

import "sort"

func ClassPhoto(redHeights, blueHeights []int) bool {
	sort.Ints(redHeights)
	sort.Ints(blueHeights)
	redHeights = reverseInPlace(redHeights)
	blueHeights = reverseInPlace(blueHeights)

	var backRow []int
	var frontRow []int
	if redHeights[0] > blueHeights[0] {
		backRow, frontRow = redHeights, blueHeights
	} else {
		backRow, frontRow = blueHeights, redHeights
	}
	for i := 0; i < len(backRow); i++ {
		if backRow[i] <= frontRow[i] {
			return false
		}
	}
	return true
}
