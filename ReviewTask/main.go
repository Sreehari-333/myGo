package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

// Inserting a node at the beginning

func (l *List) insertAtHead(val int) {

	newNode := &Node{data: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
}

// Inserting a node at the end

func (l *List) insertAtTail(val int) {

	newNode := &Node{data: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

// Deleting the first node that contain the given data

func (l *List) deleteTheFirst(data int) {

	if l.head == nil {
		fmt.Println("List is empty")
	}

	if l.head.data == data {
		l.head = l.head.next
		l.head.prev = nil
	} else if l.tail.data == data {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {

		current := l.head

		for current != nil {
			if data == current.data {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			current = current.next
		}
	}
}

// Display the elements head to tail

func (l *List) displayForward() {

	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		current := l.head

		for current != nil {
			fmt.Println(current.data)
			current = current.next
		}
	}
}

// Display the elements tail to head

func (l *List) displayBackward() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		current := l.tail

		for current != nil {
			fmt.Println(current.data)
			current = current.prev
		}
	}
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	l1 := &List{}

	// Inserting elements at the head

	fmt.Println("Enter the Element need to be insert at head")
	reader.Scan()

	head1, _ := strconv.Atoi(reader.Text())

	l1.insertAtHead(head1)

	fmt.Println("Enter the second Element need to be insert at head")
	reader.Scan()

	head2, _ := strconv.Atoi(reader.Text())

	l1.insertAtHead(head2)

	fmt.Println("Enter the third Element need to be insert at head")
	reader.Scan()

	head3, _ := strconv.Atoi(reader.Text())

	l1.insertAtHead(head3)

	// Inserting Value at tail

	fmt.Println("Enter the Element need to be insert at tail")
	reader.Scan()

	tail1, _ := strconv.Atoi(reader.Text())

	l1.insertAtTail(tail1)

	fmt.Println("Enter the second Element need to be insert at tail")
	reader.Scan()

	tail2, _ := strconv.Atoi(reader.Text())

	l1.insertAtTail(tail2)

	fmt.Println("Displaying elements from front")
	l1.displayForward()

	fmt.Println("Displaying elements from back")
	l1.displayBackward()

	// Delete the Node using value

	fmt.Println("Enter the value to delete")
	reader.Scan()

	dlt, _ := strconv.Atoi(reader.Text())
	l1.deleteTheFirst(dlt)

	fmt.Println("After Deleting ")
	l1.displayForward()

}
