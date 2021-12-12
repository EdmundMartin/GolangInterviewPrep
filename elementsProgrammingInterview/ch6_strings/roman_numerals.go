package ch6_strings

import "strings"

func RomanToInt(value string) int {
	value = strings.ToUpper(value)
	sum := 0
	nums := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}

	stringValue := strings.Split(value, "")

	for i := 0; i < len(stringValue); i++ {
		intValue := nums[stringValue[i]]

		if i+1 < len(stringValue) && nums[stringValue[i+1]] > intValue {
			sum -= intValue
		} else {
			sum += intValue
		}
	}
	return sum
}
