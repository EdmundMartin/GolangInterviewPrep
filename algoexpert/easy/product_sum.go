package easy

type SpecialArray []interface{}


func prodSumHelper(array []interface{}, depth int) int {
	prodSum := 0
	for _, val := range array {
		switch v := val.(type) {
		case int:
			prodSum += v
		case SpecialArray:
			prodSumHelper(v, depth + 1)
		}
	}
	return prodSum * depth
}

func ProductSum(array []interface{}) int {
	return prodSumHelper(array, 1)
}