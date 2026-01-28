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

func (l *LinkedList) DeleteFirst() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	l.head = l.head.next
}

func (l *LinkedList) DeleteByValue(value int) {
	// check list empty
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// if need to delete head node
	if l.head.data == value {
		l.head = l.head.next
		return
	}

	temp := l.head
	for temp.next != nil && temp.next.data != value {
		temp = temp.next
	}

	if temp.next == nil {
		fmt.Println("Value not found")
		return
	}

	temp.next = temp.next.next
}

func (l *LinkedList) Search(value int) bool {
	temp := l.head
	for temp != nil {
		if temp.data == value {
			return true
		}
		temp = temp.next
	}
	return false
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

	// delete node
	list.DeleteFirst()

	// print after delete
	fmt.Print("Print list after delete first node: ")
	list.PrintList()

	// delete node by value
	list.DeleteByValue(3)
	// print after delete
	fmt.Print("Print list after delete 3: ")
	list.PrintList()

	// search value
	isFound := list.Search(2)
	if isFound {
		fmt.Println("Value found")
	} else {
		fmt.Println("Value not found")
	}

}
