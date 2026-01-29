package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
}

func (l *Linkedlist) InsertAtBeginning(value int) {
	newNode := &Node{
		data: value,
		next: l.head,
	}

	l.head = newNode
}

func (l *Linkedlist) Reverse() {
	var prev *Node = nil
	current := l.head
	var next *Node = nil

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}

func (l *Linkedlist) Printlist() {
	temp := l.head

	for temp != nil {
		fmt.Print(temp.data, "-->")
		temp = temp.next
	}
	fmt.Print("nil")
}
func main() {
	list := Linkedlist{}

	// insert element
	list.InsertAtBeginning(5)
	list.InsertAtBeginning(6)
	list.InsertAtBeginning(7)
	list.InsertAtBeginning(8)

	// print
	fmt.Println("Before Reverse:")
	list.Printlist()

	// reverse linkedlist
	list.Reverse()

	// print
	fmt.Println("\nAfter Reverse:")
	list.Printlist()

}
