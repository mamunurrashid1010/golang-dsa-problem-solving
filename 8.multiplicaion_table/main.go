package main

import "fmt"

func main() {
	var n int

	// get n value
	fmt.Print("Enter n value: ")
	fmt.Scan(&n)

	// n multiplication table
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
