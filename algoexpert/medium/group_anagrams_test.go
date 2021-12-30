package medium

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testWords := []string{"yo", "oy", "act", "tac"}
	result := GroupAnagrams(testWords)
	fmt.Println(result)
}
