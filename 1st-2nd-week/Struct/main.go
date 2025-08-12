package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) displayDetails() {

	fmt.Printf("My name is %s . And I am %d years old", p.name, p.age)
}
func main() {

	person1 := Person{
		name: "Sreehari",
		age:  19,
	}

	person1.displayDetails()

}
