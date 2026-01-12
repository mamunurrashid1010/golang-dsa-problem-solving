package main

import "fmt"

func main() {
	var n int

	// get n value
	fmt.Print("Enter n value: ")
	fmt.Scan(&n)

	// 1 to n odd number sum calculation
	sum := 0

	// Method 1
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			sum += i
		}
	}

	// Method 2:
	for i := 1; i <= n; i += 2 {
		sum += i
	}

	fmt.Printf("1 to %d sum is: %d", n, sum)
}
