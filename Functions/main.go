package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func operations(num1, num2 int, scanner *bufio.Scanner) int {

	fmt.Println("Enter the type of operation")
	fmt.Println("1. Sum \n2. Difference \n3. Product \n4.Quotient")
	scanner.Scan()

	operation, _ := strconv.Atoi(scanner.Text())

	if operation == 1 {
		return num1 + num2
	} else if operation == 2 {
		return num1 - num2
	} else if operation == 3 {
		return num1 * num2
	} else if operation == 4 {
		return num1 / num2
	}
	return 0
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the first number")
	scanner.Scan()

	num1, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter the Second number")
	scanner.Scan()

	num2, _ := strconv.Atoi(scanner.Text())

	fmt.Println("The Result is : ", operations(num1, num2, scanner))

}
