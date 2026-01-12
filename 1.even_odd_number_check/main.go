package main

import "fmt"

func main() {
	var number int
	// get number from user
	fmt.Print("Enter you number:")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("%d is Even Number\n", number)
	} else {
		fmt.Printf("%d is Odd Number\n", number)
	}

}
