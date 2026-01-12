package main

import "fmt"

func main() {
	var n int

	// get n number
	fmt.Print("Enter n number: ")
	fmt.Scan(&n)

	// 1 to n even number print
	fmt.Printf("1 to %d even number:\n", n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
