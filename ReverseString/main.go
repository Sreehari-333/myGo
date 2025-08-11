package main

import "fmt"

func reverseString(str string) string {

	var reversedStr string
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	return reversedStr
}

func main() {

	var str string
	fmt.Println("Enter the string to reverse")
	fmt.Scan(&str)

	fmt.Println(reverseString(str))
}
