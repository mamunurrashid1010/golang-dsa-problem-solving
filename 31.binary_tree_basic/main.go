package main

import "fmt"

/*
	# Binary tree
		 10   (root)
		/  \
		5    20
	left	right
	child	child

*/

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	// add element
	root := &Node{data: 10}
	root.left = &Node{data: 5}
	root.right = &Node{data: 20}

	// print
	fmt.Println("Root: ", root.data)
	fmt.Println("Left child: ", root.left.data)
	fmt.Println("Right child: ", root.right.data)

}
