package main

import "fmt"

func main() {
	var n int
	sum := 0

	// n value input from user
	fmt.Print("Enter length of array: ")
	fmt.Scan(&n)

	// make array
	arr := make([]int, n)

	// get array value from user and sum calculation
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Enter %d value:", i+1)
		fmt.Scan(&arr[i])
		sum += arr[i]
	}

	// print array
	fmt.Println("Array value is:", arr)

	// print sum
	fmt.Println("Sum is: ", sum)

}
