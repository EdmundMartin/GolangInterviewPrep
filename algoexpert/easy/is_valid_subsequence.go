package easy


func IsValidSubsequence(array []int, sequence []int) bool {
	if len(sequence) > len(array) {
		return false
	}
	idx := 0
	for _, value := range array {
		if value == sequence[idx] {
			idx++
		}
		if idx == len(sequence) {
			return true
		}
	}
	return false
}