package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a new value to the BST
func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}

	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}

	return root
}

// GetHeight calculates the max depth of the tree
func GetHeight(root *Node) int {
	if root == nil {
		return -1
	}

	leftHeight := GetHeight(root.Left)
	rightHeight := GetHeight(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

// Delete removes a node and handles all 3 tasks (leaf, 1 child, 2 children)
func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.Value {
		root.Left = Delete(root.Left, key)
	} else if key > root.Value {
		root.Right = Delete(root.Right, key)
	} else {
		// Leaf node or node with one child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Node with two children
		// Find the smallest value in the right subtree (In-order successor)
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		// Replace current node value with successor's value
		root.Value = successor.Value
		// Delete the successor node
		root.Right = Delete(root.Right, successor.Value)
	}

	return root
}

func main() {
	var root *Node
	values := []int{50, 30, 70, 20, 40, 60, 80}

	for _, value := range values {
		root = Insert(root, value)
	}

	// Print height
	fmt.Printf("Initial Height: %d\n", GetHeight(root))

	fmt.Println("--- Deletion Tasks ---")

	// Delete a leaf (20)
	root = Delete(root, 20)
	fmt.Println("Deleted 20 (Leaf)")

	// Delete node with one child (let's say we delete 30 after 20 is gone)
	root = Delete(root, 30)
	fmt.Println("Deleted 30 (One Child)")

	// Delete node with two children (50 - Root)
	root = Delete(root, 50)
	fmt.Println("Deleted 50 (Two Children)")

	fmt.Printf("Final Height: %d\n", GetHeight(root))
}
