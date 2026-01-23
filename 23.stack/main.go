package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}

	topElement := s.data[len(s.data)-1]
	return topElement

}

func (s *Stack) isEmpty() bool {
	if len(s.data) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	stack := Stack{}

	// add element (push)
	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	stack.Push(20)

	// stack data print
	fmt.Println("Stack elements: ", stack.data)

	// remove element (pop)
	fmt.Println("Pop: ", stack.Pop())
	fmt.Println("Pop: ", stack.Pop())

	// print stack top element (peek)
	fmt.Println("Top element: ", stack.Peek())

	// is empty
	isStackEmpty := stack.isEmpty()
	if isStackEmpty {
		fmt.Println("Stack is empty")
	} else {
		fmt.Println("Stack is not empty")
	}

}
