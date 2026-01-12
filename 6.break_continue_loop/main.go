package main

import "fmt"

func main() {

	// break example
	fmt.Println("Break example: ")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	// continue example
	fmt.Println("Continue example: ")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println(i)
	}
}
