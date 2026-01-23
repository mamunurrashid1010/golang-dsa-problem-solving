package main

import "fmt"

type Stack struct {
	data []int
}

func displayMenu() {
	fmt.Println()
	fmt.Println("1.View stack elements")
	fmt.Println("2.Add stack element")
	fmt.Println("3.Remove stack element")
	fmt.Println("4.Peek stack top element")
	fmt.Println("5.Check stack is empty")
	fmt.Println("6.Exit")
}

// Print stack elements
func (s *Stack) PrintElements() {
	fmt.Println("Stack elements: ", s.data)
}

// Add element
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Remove element
func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

// Peek top element
func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}

	peekTopElement := s.data[len(s.data)-1]
	return peekTopElement
}

// stack is empty
func (s *Stack) IsEmpty() {
	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
	} else {
		fmt.Println("Stack is not empty")
	}
}
func main() {
	var choice, value int
	stack := Stack{}
	running := true

	for running {
		displayMenu()
		fmt.Print("What is your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				stack.PrintElements()
			}
		case 2:
			{
				fmt.Print("Enter value: ")
				fmt.Scan(&value)
				stack.Push(value)
			}
		case 3:
			{
				popped := stack.Pop()
				if popped != -1 {
					fmt.Println("Removed:", popped)
				}
			}
		case 4:
			{
				element := stack.Peek()
				if element != -1 {
					fmt.Println("Stack top element is: ", element)
				}
			}
		case 5:
			{
				stack.IsEmpty()
			}
		case 6:
			{
				running = false
			}
		default:
			{
				fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
			}

		}

	}

}
