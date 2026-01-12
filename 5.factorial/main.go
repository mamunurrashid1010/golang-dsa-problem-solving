package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1 //Base case
	}

	return n * factorial(n-1) // recursive call
}

func main() {
	var n int

	// get n value
	fmt.Print("Enter n value: ")
	fmt.Scan(&n)

	// Method 1: Recursion
	fact_value := factorial(n)

	// Method 2: factorial calculation using loop
	// fact := 1
	// for i := 1; i <= n; i++ {
	// 	fact *= i
	// }

	fmt.Printf("%d factorial is: %d", n, fact_value)
}
