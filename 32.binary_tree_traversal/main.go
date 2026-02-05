package main

import "fmt"

/*
	# Binary tree
		 10   (root)
		/  \
		5    20
	left	right
	child	child

	# Tree Traversal (DFS)
		1. Inorder (Left → Root → Right)
		2. Preorder (Root → Left → Right)
		3. Postorder (Left → Right → Root)

		| Traversal | Output  |
		| --------- | ------- |
		| Inorder   | 5 10 20 |
		| Preorder  | 10 5 20 |
		| Postorder | 5 20 10 |

	# Real systems used:
		* File systems (folders)
		* HTML DOM
		* Database indexing
		* AI decision trees
		* Org charts
		* JSON/XML structure

*/

type Node struct {
	data  int
	left  *Node
	right *Node
}

// Inorder (Left → Root → Right)
func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.left)
	fmt.Print(root.data, " ")
	Inorder(root.right)
}

// Preorder (Root → Left → Right)
func Preorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.data, " ")
	Inorder(root.left)
	Inorder(root.right)
}

// Postorder (Left → Right → Root)
func Postorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.left)
	Inorder(root.right)
	fmt.Print(root.data, " ")
}

func main() {
	// add element
	root := &Node{data: 10}
	root.left = &Node{data: 5}
	root.right = &Node{data: 20}

	// print inorder
	fmt.Print("Inorder: ")
	Inorder(root)

	// print preorder
	fmt.Print("\nPreorder: ")
	Preorder(root)

	// print postorder
	fmt.Print("\nPostorder: ")
	Postorder(root)

}
