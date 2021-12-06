package ch3_stacks_queues

type SimpleStack struct {
	Stack []int
}

func NewSimpleStack() *SimpleStack {
	return &SimpleStack{Stack: []int{}}
}

func (s *SimpleStack) Push(value int) {
	s.Stack = append(s.Stack, value)
}

func (s *SimpleStack) Pop() int {
	var value int
	value, s.Stack = s.Stack[len(s.Stack)-1], s.Stack[:len(s.Stack)-1]
	return value
}

func (s *SimpleStack) Peek() int {
	return s.Stack[len(s.Stack)-1]
}

func (s *SimpleStack) Length() int {
	return len(s.Stack)
}



func SortStack(stack *SimpleStack) *SimpleStack {
	replacement := NewSimpleStack()
	for stack.Length() > 0 {
		tmp := stack.Pop()
		for replacement.Length() > 0 && replacement.Peek() > tmp {
			stack.Push(replacement.Pop())
		}
		replacement.Push(tmp)
	}

	for replacement.Length() > 0 {
		stack.Push(replacement.Pop())
	}

	return stack
}