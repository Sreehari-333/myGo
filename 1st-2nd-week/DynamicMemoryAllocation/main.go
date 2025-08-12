package main

import "fmt"

func main() {

	p := new(int)
	*p = 100
	fmt.Println(*p)

	slice := make([]int, 3)

	slice[0] = 10
	slice[1] = 20
	slice[2] = 30

	fmt.Println(slice)
}
