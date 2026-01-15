package main

import "fmt"

func main() {
	target := 5

	// slice define
	arr := []int{1, 5, 2, 3, 5, 5, 7, 8, 5}

	// target total found count
	targetCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			targetCount++
		}
	}

	// print
	fmt.Printf("%d found total: %d", target, targetCount)
}
