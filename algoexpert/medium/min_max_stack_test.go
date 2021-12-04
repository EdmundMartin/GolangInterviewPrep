package medium

import (
	"fmt"
	"testing"
)

func TestMinMaxStack(t *testing.T) {
	stack := NewMinMaxStack()
	stack.Push(10)
	stack.Push(8)
	stack.Push(12)
	fmt.Println(stack.String())
	stack.Pop()
	fmt.Println(stack.String())
}