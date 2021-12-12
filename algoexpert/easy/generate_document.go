package easy

func GenerateDocument(characters, document string) bool {
	charMap := make(map[rune]int)
	for _, char := range characters {
		_, ok := charMap[char]
		if !ok {
			charMap[char] = 1
		} else {
			charMap[char]++
		}
	}
	for _, char := range document {
		val, ok := charMap[char]
		if !ok {
			return false
		}
		if val == 0 {
			return false
		}
		charMap[char]--
	}
	return true
}
