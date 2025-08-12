package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

// To add a node in the beginning

func (list *LinkedList) insertFront(val int) {

	newNode := &Node{data: val}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {

		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}
}

// To delete the last node

func (list *LinkedList) deleteLast() {

	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}
}

// To display

func (list *LinkedList) display() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	l1 := &LinkedList{}

	l1.insertFront(10)
	l1.insertFront(20)

	l1.deleteLast()

	l1.display()

}
