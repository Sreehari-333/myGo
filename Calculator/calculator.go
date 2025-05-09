package main

import "fmt"

func main() {
	var num1, num2 int
	var operation string

	fmt.Printf("Enter Two Numbers \n")
	fmt.Scan(&num1, &num2)
	fmt.Printf("Enter the operation \n")
	fmt.Scan(&operation)

	switch operation {
	case "add":
		fmt.Printf("%d + %d = %d", num1, num2, num1+num2)
	case "subtract":
		fmt.Printf("%d - %d = %d", num1, num2, num1-num2)
	case "multiply":
		fmt.Printf("%d * %d = %d", num1, num2, num1*num2)
	case "divide":
		fmt.Printf("%d / %d = %d", num1, num2, num1/num2)
	default:
		fmt.Printf("Invalid Operation")
	}

}
