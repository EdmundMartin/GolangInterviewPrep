package medium

import "strings"

type SimpleStringSet struct {
	values map[string]interface{}
}

func NewSimpleStringSet(values ...string) *SimpleStringSet {
	set := &SimpleStringSet{values: make(map[string]interface{})}
	for _, val := range values {
		set.values[val] = nil
	}
	return set
}

func (s SimpleStringSet) Contains(val string) bool {
	_, ok := s.values[val]
	return ok
}

func BalancedBrackets(input string) bool {
	openingSet := NewSimpleStringSet("(", "{", "[")
	closingSet := NewSimpleStringSet(")", "}", "]")
	matchingBrackets := map[string]string{
		"}":"{",
		")":"(",
		"]":"[",
	}
	var stack []string
	for _, char := range strings.Split(input, "") {
		if openingSet.Contains(char) {
			stack = append(stack, char)
		} else if closingSet.Contains(char) {
			if len(stack) == 0 {
				return false
			}
			val, _ := matchingBrackets[char]
			if stack[len(stack) - 1] == val {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}