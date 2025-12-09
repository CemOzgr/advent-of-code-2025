package day3

type Stack struct {
	top int
	arr []int
}

func NewStack(size int) *Stack {
	return &Stack{arr: make([]int, size)}
}

func (stack *Stack) Pop() int {
	if stack.top == 0 {
		panic("empty pop")
	}

	stack.top--
	return stack.arr[stack.top]
}

func (stack *Stack) Push(value int) {
	if stack.top == len(stack.arr) {
		panic("stack full")
	}

	stack.arr[stack.top] = value
	stack.top++
}

func (stack *Stack) Peek() int {
	return stack.arr[stack.top-1]
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == 0
}

func (stack *Stack) ShouldPop(value int) bool {
	return !stack.IsEmpty() && value > stack.Peek()
}
