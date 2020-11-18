package main

import "fmt"

type Queue struct {
	/* Queue Structure. */
	front    int
	rear     int
	size     int
	capacity int
	array    [10000]*int
}

func createQueue(capacity int) *Queue {
	/* Creating a new empty Queue */
	queue := Queue{}
	queue.capacity = capacity
	queue.rear = capacity - 1

	return &queue
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func (q *Queue) isFull() bool {
	return q.size == q.capacity
}

func (q *Queue) enqueue(data int) {
	/* Inserting a element in rear of the queue.  */
	if q.isFull() {
		return
	}

	q.rear = (q.rear + 1) % q.capacity
	q.array[q.rear] = &data
	q.size = q.size + 1
	fmt.Printf("\n%d is enqueued to queue\n", data)
}

func (q *Queue) dequeue() int {
	/* Deleting a element from the queue. */
	if q.isEmpty() {
		return 0
	}

	var data = q.array[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size = q.size - 1
	fmt.Printf("\n%d is dequeued from queue\n", *data)
	return *data
}

func (q *Queue) getFront() int {
	/* Get the first element of the queue. */
	if q.isEmpty() {
		return 0
	}

	return *q.array[q.front]
}

func (q *Queue) getRear() int {
	/* Get the last element of the queue. */
	if q.isEmpty() {
		return 0
	}

	return *q.array[q.rear]
}

func main() {
	queue := createQueue(100) // Creating a new Queue.

	fmt.Printf("\nIs the queue empty: %v\n", queue.isEmpty())
	queue.enqueue(5)
	queue.enqueue(84)
	queue.enqueue(3)
	queue.enqueue(6)
	queue.enqueue(45)
	queue.enqueue(9)

	queue.dequeue()
	queue.dequeue()

	fmt.Printf("\nIs the queue full: %v\n", queue.isFull())
	fmt.Printf("\nFront element: %v\n", queue.getFront())
	fmt.Printf("\nRear element: %v\n", queue.getRear())
}
