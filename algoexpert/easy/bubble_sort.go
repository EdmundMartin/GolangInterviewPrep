package easy

func BubbleSort(array []int) []int {
	loops := len(array)
	for loops > 0 {
		swaps := 0
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swaps++
			}
		}
		if swaps == 0 {
			break
		}
		loops--
	}
	return array
}
