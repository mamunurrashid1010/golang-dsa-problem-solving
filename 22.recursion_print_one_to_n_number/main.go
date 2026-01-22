package main

import "fmt"

func printNumber(n int) {
	if n == 0 {
		return // base case
	}

	printNumber(n - 1) // recursive call
	fmt.Println(n)
}

func main() {
	var n int

	// get n value
	fmt.Print("Enter n value: ")
	fmt.Scan(&n)

	// print 1 to n using recursion
	printNumber(n)
}

// What is Base Case and why it is important?
// Base case is so important to stop iteration. If we not use base case the program goes to infinit loop.