package ch1_arrays_strings

func oneEditReplace(first, second string) bool {
	foundDiff := false
	for idx, _ := range first {
		if first[idx] != second[idx] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}
	return true
}

func oneEditInsert(first, second string) bool {
	firstIdx, secondIdx := 0, 0
	for secondIdx < len(second) && firstIdx < len(first) {
		if first[firstIdx] != second[secondIdx] {
			if firstIdx != secondIdx {
				return false
			}
			secondIdx++
		} else {
			firstIdx++
			secondIdx++
		}
	}
	return true
}

func OneEditAway(first string, second string) bool {
	if len(first) == len(second) {
		return oneEditReplace(first, second)
	} else if len(first) + 1 == len(second) {
		return oneEditInsert(first, second)
	} else if len(first) - 1 == len(second) {
		return oneEditInsert(second, first)
	}
	return false
}