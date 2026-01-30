package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

func (l *CircularLinkedList) InsertAtBeginning(value int) {
	newNode := &Node{data: value}

	// if list is empty
	if l.head == nil {
		l.head = newNode
		newNode.next = l.head
		return
	}

	temp := l.head
	for temp.next != l.head {
		temp = temp.next
	}

	temp.next = newNode
	newNode.next = l.head
	l.head = newNode

}

func (l *CircularLinkedList) PrintList() {

	if l.head == nil {
		fmt.Println("List empty")
		return
	}

	temp := l.head
	for {
		fmt.Print(temp.data, "-->")
		temp = temp.next
		if temp == l.head {
			break
		}
	}
	fmt.Println("(back to head)")
}

func main() {
	circularList := CircularLinkedList{}

	// insert element
	circularList.InsertAtBeginning(5)
	circularList.InsertAtBeginning(10)
	circularList.InsertAtBeginning(15)

	// print
	circularList.PrintList()

}
