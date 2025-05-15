package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	Head *Node
}

// Adding values

func (list *linkedList) insert(data int) {

	newNode := &Node{data: data}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head

	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

// Delete the node by value

func (list *linkedList) deleteByValue(value int) {

	if list.Head.data == value {
		list.Head = list.Head.next
		return
	}

	current := list.Head

	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// Display the value

func (list linkedList) display() {

	current := list.Head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// Searching the values

func (list linkedList) search(value int) bool {

	current := list.Head

	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	l1 := &linkedList{}
	arr := []int{10, 20, 30, 40, 50, 60}
	for _, num := range arr {
		l1.insert(num)
	}

	l1.deleteByValue(20)
	l1.display()
	fmt.Println(l1.search(20))
}
