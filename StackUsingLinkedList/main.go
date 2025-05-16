package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(val int) {

	newNode := &Node{data: val}

	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) pop() int {

	if s.top == nil {
		fmt.Println("Stack is empty")
	}

	val := s.top.data
	s.top.next = s.top
	return val
}

func (s *Stack) peek() {
	fmt.Println(s.top.data)
}

func main() {
	myStack := &Stack{}

	myStack.push(10)
	myStack.push(20)
	myStack.push(30)

	fmt.Println(myStack.pop())

	myStack.peek()
}
