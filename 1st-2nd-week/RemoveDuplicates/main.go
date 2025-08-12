package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2, 3, 4, 4, 5}
	uniqueNumbers := []int{}
	seen := make(map[int]bool)

	for _, num := range numbers {
		if !seen[num] {
			seen[num] = true

			uniqueNumbers = append(uniqueNumbers, num)
		}
	}

	fmt.Println(uniqueNumbers)
}
