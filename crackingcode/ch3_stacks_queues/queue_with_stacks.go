package ch3_stacks_queues

import "errors"

// Implement an instance of MyQueue which implements a queue using two stacks.

type MyQueue struct {
	oldest []int
	newest []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		oldest: []int{},
		newest: []int{},
	}
}

func (m *MyQueue) add(value int) {
	m.newest = append(m.newest, value)
}

func (m *MyQueue) Peek() (int, error) {
	m.shiftStacks()
	if len(m.oldest) == 0 {
		return 0, errors.New("empty stack")
	}
	return m.oldest[len(m.oldest) - 1], nil
}

func (m *MyQueue) Pop() (int, error) {
	m.shiftStacks()
	if len(m.oldest) == 0 {
		return 0, errors.New("empty stack")
	}
	var x int
	x, m.newest = m.newest[len(m.newest)-1], m.newest[:len(m.newest)-1]
	return x, nil
}

func (m *MyQueue) shiftStacks() {
	if len(m.oldest) == 0 {
		for len(m.newest) > 0 {
			var x int
			x, m.newest = m.newest[len(m.newest)-1], m.newest[:len(m.newest)-1]
			m.oldest = append(m.oldest, x)
		}
	}
}