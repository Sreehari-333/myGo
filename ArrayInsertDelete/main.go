package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func insertion(arr []int, scanner *bufio.Scanner) []int {

	fmt.Println("Enter the value to insert")
	scanner.Scan()

	value, _ := strconv.Atoi(scanner.Text())

	arr = append(arr, value)

	return arr

}

func deletion(arr []int, scanner *bufio.Scanner) []int {

	fmt.Println("Enter the index to delete ")
	scanner.Scan()

	index, _ := strconv.Atoi(scanner.Text())

	arr = append(arr[:index], arr[index+1:]...)

	return arr
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	arr := []int{1, 3, 4, 3, 5, 6}

	fmt.Println(insertion(arr, scanner))

	fmt.Println(deletion(arr, scanner))

}
