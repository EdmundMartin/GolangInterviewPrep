package medium

import "fmt"

type MinMax struct {
	Min int
	Max int
}

type MinMaxStack struct {
	StackMinMax []MinMax
	Stack []int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{
		StackMinMax: []MinMax{},
		Stack:       []int{},
	}
}

func (m *MinMaxStack) String() string {
	return fmt.Sprintf("Stack: %v, StackMinMax: %v", m.Stack, m.StackMinMax)
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m *MinMaxStack) Peek() int {
	return m.Stack[len(m.Stack) - 1]
}

func (m *MinMaxStack) Pop() int {
	var x int
	x, m.Stack = m.Stack[len(m.Stack) - 1], m.Stack[:len(m.Stack) - 1]
	m.StackMinMax = m.StackMinMax[:len(m.StackMinMax) - 1]
	return x
}

func (m *MinMaxStack) Push(value int) {
	newMinMax := MinMax{
		Min: value,
		Max: value,
	}
	if len(m.StackMinMax) > 0 {
		lastMinMax := m.StackMinMax[len(m.StackMinMax) - 1]
		newMinMax.Min = MinInt(lastMinMax.Min, newMinMax.Min)
		newMinMax.Max = MaxInt(lastMinMax.Max, newMinMax.Max)
	}
	m.StackMinMax = append(m.StackMinMax, newMinMax)
	m.Stack = append(m.Stack, value)
}

func (m *MinMaxStack) GetMin() int {
	return m.StackMinMax[len(m.StackMinMax) - 1].Min
}

func (m *MinMaxStack) GetMax() int {
	return m.StackMinMax[len(m.StackMinMax) - 1].Max
}