package main

import "fmt"

func main() {

	var num int

	fmt.Println("Enter the Number wants to check")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Printf("%d is an even number", num)
	} else {
		fmt.Printf("%d is an odd number", num)
	}
}
