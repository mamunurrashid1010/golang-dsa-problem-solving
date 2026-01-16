package main

import "fmt"

func isSorted(arr []int) bool {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true

}
func main() {
	// slice declare
	arr := []int{1, 5, 3, 2, 1}

	// check array asc sorted or not
	isSorted := isSorted(arr)

	if isSorted {
		fmt.Println("Array is sorted.")
	} else {
		fmt.Println("Array is not sorted")
	}
}
