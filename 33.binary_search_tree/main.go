/*
# Binary Search Tree (BST)
A Binary Search Tree (BST) is a hierarchical data structure that organizes elements in a way that allows for efficient search, insertion, and deletion operations.
It is a special type of binary tree, meaning each node has at most two children (left and right), and it adheres to a specific ordering property.

        10
       /  \
      5    20
     / \     \
    2   7     30

# rule
Left Subtree  <  Root  <  Right Subtree

# BST Properties
Left child < Parent
Right child > Parent
Get sorted output- Inorder traversal use

*/

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{data: value}
	}

	if value < root.data {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	return root

}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Print(root.data, " ")
	inorder(root.right)
}

func search(root *Node, target int) bool {
	if root == nil {
		return false
	}

	if root.data == target {
		return true
	} else if target < root.data {
		return search(root.left, target)
	} else {
		return search(root.right, target)
	}
}

func minimun(root *Node) int {
	if root == nil {
		return -1
	}

	if root.left == nil {
		return root.data
	}

	return minimun(root.left)

}

func maximun(root *Node) int {
	if root == nil {
		return -1
	}

	if root.right == nil {
		return root.data
	}

	return maximun(root.right)
}

func countNodes(root *Node) int {
	count := 0
	if root == nil {
		return count
	}

	return 1 + countNodes(root.left) + countNodes(root.right)
}

func main() {
	var root *Node

	values := []int{10, 5, 20, 7, 2, 30}

	// add node
	for _, value := range values {
		root = insert(root, value)
	}

	// print: inorder traversal
	fmt.Print("Inorder Traversal (Sorted): ")
	inorder(root)

	// search
	fmt.Print("\nSearch 20: ", search(root, 20))
	fmt.Print("\nSearch 15: ", search(root, 15))

	// minimun value
	minValue := minimun(root)
	if minValue == -1 {
		fmt.Println("Tree is empty")
	} else {
		fmt.Println("\nMinmum value: ", minValue)
	}

	// maximun value
	maxValue := maximun(root)
	if maxValue == -1 {
		fmt.Println("Tree is empty")
	} else {
		fmt.Println("Maximum value: ", maxValue)
	}

	// total node count
	fmt.Println("Total node: ", countNodes(root))

}
