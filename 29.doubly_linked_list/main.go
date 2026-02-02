package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (l *DoublyLinkedList) InsertAtBeginning(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

func (l *DoublyLinkedList) InsertAtEnd(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	newNode.prev = temp
}

func (l *DoublyLinkedList) PrintForward() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " <-> ")
		temp = temp.next
	}
	fmt.Println("nil")
}

func (l *DoublyLinkedList) PrintBackward() {

	if l.head == nil {
		fmt.Println("List empty")
		return
	}

	temp := l.head

	// go to last node
	for temp.next != nil {
		temp = temp.next
	}

	for temp != nil {
		fmt.Print(temp.data, " <-> ")
		temp = temp.prev
	}
	fmt.Println("nil")
}

func main() {
	list := DoublyLinkedList{}

	// insert element at beginning
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(5)
	list.InsertAtBeginning(3)

	// insert element at beginning
	list.InsertAtEnd(15)
	list.InsertAtEnd(20)

	// print forward
	fmt.Print("Print Forward\n")
	list.PrintForward()

	// print backward
	fmt.Print("\nPrint Backward\n")
	list.PrintBackward()

}
