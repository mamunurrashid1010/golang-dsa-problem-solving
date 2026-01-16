package main

import "fmt"

func bubbleSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		swapped := false

		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}

	}

	return arr
}

func main() {
	// slice declare
	arr := []int{2, 1, 5, 11, 10, 6}

	// before sort
	fmt.Println("Data before sort: ", arr)

	// data sorting using bubble sort
	sortedArray := bubbleSort(arr)

	// after sort
	fmt.Println("Data after sort: ", sortedArray)
}
