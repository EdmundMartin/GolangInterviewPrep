package ch3_stacks_queues

import "errors"

const numberStacks = 3

type FixedMultiStack struct {
	stackCapacity int
	values        []int
	sizes         []int
}

func NewFixedMultiStack(stackSize int) *FixedMultiStack {
	return &FixedMultiStack{
		stackCapacity: stackSize,
		values:        make([]int, stackSize*numberStacks),
		sizes:         make([]int, numberStacks),
	}
}

func (f *FixedMultiStack) Push(stackNumber, value int) error {
	if f.isFull(stackNumber) {
		return errors.New("stack is full")
	}
	f.sizes[stackNumber]++
	f.values[f.topIndex(stackNumber)] = value
	return nil
}

func (f *FixedMultiStack) Pop(stackNumber int) (int, error) {
	if f.isEmpty(stackNumber) {
		return 0, errors.New("stack is empty")
	}
	topIdx := f.topIndex(stackNumber)
	value := f.values[topIdx]
	f.values[topIdx] = 0
	f.sizes[stackNumber]--
	return value, nil
}

func (f *FixedMultiStack) Peek(stackNumber int) (int, error) {
	if f.isEmpty(stackNumber) {
		return 0, errors.New("stack is empty")
	}
	value := f.values[f.topIndex(stackNumber)]
	return value, nil
}

func (f *FixedMultiStack) isEmpty(stackNumber int) bool {
	return f.sizes[stackNumber] == 0
}

func (f *FixedMultiStack) isFull(stackNumber int) bool {
	return f.sizes[stackNumber] == f.stackCapacity
}

func (f *FixedMultiStack) topIndex(stackNumber int) int {
	offset := stackNumber * f.stackCapacity
	size := f.sizes[stackNumber]
	return offset + size - 1
}
