package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	target := 9

	// slice define
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// get target index
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Println("Target found at index: ", index)
	} else {
		fmt.Println("Target not found.")
	}
}
