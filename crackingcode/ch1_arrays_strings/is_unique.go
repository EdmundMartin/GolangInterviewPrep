package ch1_arrays_strings

/*
Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use
additional data structures?

It's okay to assume 256 characters. This would be the case in extended ASCII. You should clarify your assumptions
with your interview.
*/

func IsUnique(target string) bool {
	if len(target) > 128 {
		return false
	}
	charSet := [128]bool{}
	for _, charRune := range target {
		value := int(charRune)
		exists := charSet[value]
		if exists {
			return false
		}
		charSet[value] = true
	}
	return true
}
