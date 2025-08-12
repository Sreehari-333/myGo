package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) pop() {
	if len(s.data) > 0 {
		s.data = s.data[:len(s.data)-1]
	}
}

func (s *Stack) peek() int {

	if len(s.data) > 0 {
		last := s.data[len(s.data)-1]
		return last
	}
	return 0
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func main() {

	myStack := &Stack{}

	myStack.push(10)
	myStack.push(20)
	myStack.push(30)

	myStack.pop()

	for _, val := range myStack.data {
		fmt.Println(val)
	}

	fmt.Println("Last value is :", myStack.peek())
	fmt.Println("Stack is :", myStack.isEmpty())
}
