package ch1_arrays_strings

func buildFreqTable(phrase string) map[int]int {
	counter := make(map[int]int)
	for _, char := range phrase {
		val := int(char)
		if val == 32 {
			continue
		}
		_, ok := counter[val]
		if ok {
			counter[val]++
		} else {
			counter[val] = 1
		}
	}
	return counter
}

func checkMaxOneOdd(counter map[int]int) bool {
	foundOdd := false
	for _, v := range counter {
		if v%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}

func PalindromePermutationHashTable(phrase string) bool {
	counter := buildFreqTable(phrase)
	return checkMaxOneOdd(counter)
}

func getCharValue(r rune) int {
	a := int('a')
	z := int('z')
	val := int(r)
	if a <= val && val <= z {
		return val - a
	}
	return -1
}

func PalindromePermutationOptimised(phrase string) bool {
	countOdd := 0
	size := int('z') - int('a') + 1
	table := make([]int, size)
	for _, char := range phrase {
		val := getCharValue(char)
		if val != -1 {
			table[val]++
			if table[val]%2 == 1 {
				countOdd++
			} else {
				countOdd--
			}
		}
	}
	return countOdd <= 1
}
