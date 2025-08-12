package main

import "fmt"

func main() {

	numbers := []int{10, 5, 8, 20, 2, 18}

	max := numbers[0]
	min := numbers[0]

	for _, num := range numbers {

		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}

	fmt.Println("Largest is : ", max)
	fmt.Println("Smallest is : ", min)
}
