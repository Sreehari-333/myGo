package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

func (q *Queue) enqueue(val int) {

	newNode := &Node{value: val}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *Queue) dequeue() int {

	removed := q.front.value
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}
	return removed
}

func main() {

	myQueue := &Queue{}

	myQueue.enqueue(10)
	myQueue.enqueue(20)
	myQueue.enqueue(20)

	fmt.Println(myQueue.front.value)
	fmt.Println(myQueue.front.next.value)

	fmt.Println("Removed value :", myQueue.dequeue())

}
