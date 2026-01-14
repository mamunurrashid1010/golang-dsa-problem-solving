package main

import "fmt"

func main() {
	var n int

	// array length input from user
	fmt.Print("Enter length of array: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Array is empty")
		return
	}

	// array make
	arr := make([]int, n)

	// array value input from user
	for i := 0; i < n; i++ {
		fmt.Printf("Enter %d value: ", i)
		fmt.Scan(&arr[i])
	}

	// default value set
	largest := arr[0]
	smallest := arr[0]

	// largest, smallest value find
	for i := 1; i < n; i++ {
		if arr[i] > largest {
			largest = arr[i]
		} else if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	// print
	fmt.Println("The largest value is: ", largest)
	fmt.Println("The smallest value is: ", smallest)

}
