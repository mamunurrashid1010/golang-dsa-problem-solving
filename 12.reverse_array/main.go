package main

import "fmt"

func main() {
	// array
	arr := [5]int{1, 2, 3, 4, 5}

	// before reverse
	fmt.Println("Array before reverse: ", arr)

	// array pointer set
	l := 0
	r := len(arr) - 1

	// reverse array
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}

	// after reverse
	fmt.Println("Array after reverse: ", arr)

}
