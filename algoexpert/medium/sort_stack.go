package medium

type SimpleStack struct {
	Stack []int
}


func NewStack(values... int) *SimpleStack {
	s := &SimpleStack{Stack: []int{}}
	for _, val := range values {
		s.Append(val)
	}
	return s
}


func (s *SimpleStack) Pop() int {
	var x int
	x, s.Stack = s.Stack[len(s.Stack)-1], s.Stack[:len(s.Stack)-1]
	return x
}

func (s *SimpleStack) Peek() int {
	return s.Stack[len(s.Stack) - 1]
}

func (s *SimpleStack) Append(value int) {
	s.Stack = append(s.Stack, value)
}

func (s *SimpleStack) Length() int {
	return len(s.Stack)
}


func SortStack(stack *SimpleStack) *SimpleStack {
	if stack.Length() == 0 {
		return stack
	}
	top := stack.Pop()

	SortStack(stack)

	insertSortedOrder(stack, top)

	return stack
}

func insertSortedOrder(stack *SimpleStack, value int) {
	if stack.Length() == 0 || stack.Peek() <= value {
		stack.Append(value)
		return
	}

	top := stack.Pop()

	insertSortedOrder(stack, value)

	stack.Append(top)
}