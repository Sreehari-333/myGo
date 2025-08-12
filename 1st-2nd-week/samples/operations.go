package samples

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Operations() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the first number")
	scanner.Scan()

	num1, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter the second number")
	scanner.Scan()

	num2, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter the Operation")
	scanner.Scan()

	operation := scanner.Text()

	switch operation {
	case "add":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "subtract":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "multiply":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "divide":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	}
}
