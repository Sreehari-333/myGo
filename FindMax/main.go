package main

import "fmt"

func main() {

	arr := []int{10, 2, 4, 4, 8, 110}
	max := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println("Max number in array is : ", max)
	fmt.Println("Time complexity of this program is O(n) (Linear Time) ")
}
