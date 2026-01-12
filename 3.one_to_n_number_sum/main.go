package main

import "fmt"

func main() {
	var n, sum int

	// get n number
	fmt.Print("Enter n number: ")
	fmt.Scan(&n)

	// sum calculation using Gauss formula
	sum = n * (n + 1) / 2

	// sum calculation using loop
	// for i := 1; i <= n; i++ {
	// 	sum += i
	// }

	// total
	fmt.Println("Sum = ", sum)

}
