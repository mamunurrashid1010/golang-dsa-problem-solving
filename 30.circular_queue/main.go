package main

import "fmt"

type CircularQueue struct {
	data  []int
	front int
	rear  int
	size  int
}

func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, size),
		front: -1,
		rear:  -1,
		size:  size,
	}
}

func (q *CircularQueue) IsFull() bool {
	return (q.rear+1)%q.size == q.front
}

func (q *CircularQueue) IsEmpty() bool {
	return q.front == -1
}

func (q *CircularQueue) Enqueue(value int) {
	if q.IsFull() {
		fmt.Println("Queue is full")
		return
	}

	if q.IsEmpty() {
		q.front = 0
	}

	q.rear = (q.rear + 1) % q.size
	q.data[q.rear] = value
}

func (q *CircularQueue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}

	value := q.data[q.front]

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % q.size
	}

	return value
}

func (q *CircularQueue) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	i := q.front
	for {
		fmt.Print(q.data[i], " ")
		if i == q.rear {
			break
		}
		i = (i + 1) % q.size
	}

	fmt.Println()

}

func main() {
	queue := NewCircularQueue(5)

	// insert element
	queue.Enqueue(5)
	queue.Enqueue(10)
	queue.Enqueue(15)

	// print
	fmt.Println("Queue after add element")
	queue.PrintQueue()

	// // remove element
	value := queue.Dequeue()
	if value != -1 {
		fmt.Println("Remove queue element: ", value)
	}

	// // print
	fmt.Println("Queue after remove element")
	queue.PrintQueue()
}
