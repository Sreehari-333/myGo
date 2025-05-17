package main

import "fmt"

type Queue struct {
	items []int
}

// To add the element to the back

func (q *Queue) enqueue(item int) {

	q.items = append(q.items, item)
}

// To remove the element drom front

func (q *Queue) dequeue() int {

	front := q.items[0]
	q.items = q.items[1:]
	return front
}

func main() {
	myQueue := Queue{}

	myQueue.enqueue(10)
	myQueue.enqueue(20)
	myQueue.enqueue(30)

	fmt.Println(myQueue.items)

	fmt.Println(myQueue.dequeue())

	fmt.Println(myQueue.items)

}
