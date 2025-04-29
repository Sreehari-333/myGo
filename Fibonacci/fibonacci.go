package main

import (
	"fmt"
)

func main() {
	var n int
	a, b := 0, 1
	fmt.Printf("Enter the Number \n")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
}
