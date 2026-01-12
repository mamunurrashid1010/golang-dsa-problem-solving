package main

import "fmt"

func main() {
	var number int

	// get number from user
	fmt.Print("Enter you number: ")
	fmt.Scan(&number)

	// check number is positive, negative, zero
	if number > 0 {
		fmt.Print("Number is Positive")
	} else if number < 0 {
		fmt.Print("Number is Negative")
	} else {
		fmt.Print("This is not a number")
	}
}
