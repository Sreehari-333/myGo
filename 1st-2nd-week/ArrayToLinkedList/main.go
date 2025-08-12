package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	Head *Node
}

func (list *linkedList) convert(arr []int) {

	for _, val := range arr {
		newNode := &Node{data: val}

		if list.Head == nil {
			list.Head = newNode
		} else {
			current := list.Head

			for current.next != nil {
				current = current.next
			}
			current.next = newNode
		}
	}
}

func (list linkedList) display() {

	current := list.Head

	if current == nil {
		fmt.Println("No values ")
		return
	}

	for current.next != nil {
		fmt.Println(current.data)
		current = current.next
	}

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	l1 := &linkedList{}
	l1.convert(arr)

	l1.display()
}
