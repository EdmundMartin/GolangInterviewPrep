package ch3_stacks_queues

import (
	"errors"
	"math"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Placeholder struct {
	Value int
	Min   int
}

type StackWithMin struct {
	stack []Placeholder
}

func NewStackWithMin() *StackWithMin {
	return &StackWithMin{
		stack: []Placeholder{},
	}
}

func (s *StackWithMin) Push(value int) {
	var oldMin int
	if len(s.stack) == 0 {
		oldMin = math.MaxInt32
	} else {
		oldMin = s.stack[len(s.stack)-1].Min
	}
	s.stack = append(s.stack, Placeholder{
		Value: value,
		Min:   MinInt(oldMin, value),
	})
}

func (s *StackWithMin) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack is empty")
	}
	var v Placeholder
	v, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return v.Value, nil
}

func (s *StackWithMin) Peek() int {
	return s.stack[len(s.stack)-1].Value
}

func (s *StackWithMin) Min() int {
	return s.stack[len(s.stack)-1].Min
}
