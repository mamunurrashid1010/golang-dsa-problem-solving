package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{
		data: value,
		next: l.head,
	}
	l.head = newNode
}

func (l *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{data: value}

	// is empty check
	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode

}

func (l *LinkedList) PrintList() {
	temp := l.head

	for temp != nil {
		fmt.Print(temp.data, "--->")
		temp = temp.next
	}

	fmt.Println("nil")
}

func main() {
	list := LinkedList{}

	// insert element at the Beginning
	list.InsertAtBeginning(5)
	list.InsertAtBeginning(7)
	list.InsertAtBeginning(10)

	// insert element at the end
	list.InsertAtEnd(3)
	list.InsertAtEnd(2)

	// print
	list.PrintList()

}
