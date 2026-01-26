package main

import "fmt"

type Queue struct {
	data []int
}

// Add element
func (q *Queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

// Remoe element
func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}

	front := q.data[0]
	q.data = q.data[1:]
	return front
}

// Get front element
func (q *Queue) Front() int {
	if len(q.data) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}

	return q.data[0]
}

// Queue empty check
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func main() {
	queue := Queue{}

	// add element
	queue.Enqueue(5)
	queue.Enqueue(7)
	queue.Enqueue(10)
	queue.Enqueue(20)

	// print queue
	fmt.Println("Queue: ", queue.data)

	// remove element
	queue.Dequeue()
	queue.Dequeue()

	// print queue
	fmt.Println("Queue after remove element: ", queue.data)

	// print front
	fmt.Println("Queue front element: ", queue.Front())

	// queue empty check
	isEmpty := queue.IsEmpty()
	if isEmpty {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Queue is not empty")
	}

}
