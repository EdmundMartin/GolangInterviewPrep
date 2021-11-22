package easy


func TwoSum(array []int, targetSum int) []int {
	seen := make(map[int]interface{})
	for _, num := range array {
		other := targetSum - num
		if _, ok := seen[other]; ok {
			return []int{other, num}
		}
		seen[num] = nil
	}
	return []int{}
}