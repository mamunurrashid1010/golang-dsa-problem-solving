package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1 //Base case
	}

	return n * factorial(n-1)
}

func main() {
	var n int

	// get n value form user
	fmt.Print("Enter value:")
	fmt.Scan(&n)

	// get fact value
	fact_value := factorial(n)

	fmt.Printf("%d factorial is: %d", n, fact_value)
}
