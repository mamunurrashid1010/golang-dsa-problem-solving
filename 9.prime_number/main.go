package main

import "fmt"

func primeNumberCheck(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			fmt.Printf("%d is not a prime number.", number)
			return false
		}
	}

	fmt.Printf("%d is a prime number.", number)
	return true
}

func main() {
	var number int

	// get n value
	fmt.Print("Enter Number: ")
	fmt.Scan(&number)

	// prime check
	if number < 0 {
		fmt.Println("Please enter positive number.")
	} else if number == 1 {
		fmt.Println("Please enter number greater than 1")
	} else if number > 1 {
		primeNumberCheck(number)
	} else {
		fmt.Println("Please enter a valid positive number.")
	}
}
