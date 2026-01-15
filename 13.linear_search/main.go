package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}

func main() {
	var n, target int

	// get array length
	fmt.Print("Enter length of array: ")
	fmt.Scan(&n)

	// get target
	fmt.Print("Enter target value: ")
	fmt.Scan(&target)

	// make array
	arr := make([]int, n)

	// get array element form user
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Enter value %d = ", i)
		fmt.Scan(&arr[i])
	}

	// get element index
	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Println("Target found at index:", index)
	} else {
		fmt.Println("Target not found.")
	}

}
