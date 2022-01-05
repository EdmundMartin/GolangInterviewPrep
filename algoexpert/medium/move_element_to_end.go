package medium


func MoveElementToEnd(array []int, target int) []int {
	i := 0
	j := len(array) - 1
	for i < j {
		if array[i] == target && array[j] != target {
			array[i], array[j] = array[j], array[i]
		}
		if array[i] != target {
			i++
		}
		if array[j] == target {
			j--
		}
	}
	return array
}