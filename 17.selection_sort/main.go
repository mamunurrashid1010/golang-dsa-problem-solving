package main

import "fmt"

func selectionSort(arr []int) []int {
	len := len(arr)

	for i := 0; i < len-1; i++ {
		minIndex := i

		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

func main() {
	// slice declare
	arr := []int{10, 8, 5, 15, 12, 3}

	// before sort
	fmt.Println("Data before sort: ", arr)

	// data sorting using selection sort
	sortedArray := selectionSort(arr)

	// after sort
	fmt.Println("Data after sort: ", sortedArray)
}
