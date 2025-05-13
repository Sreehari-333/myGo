package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}

	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	fmt.Println("Reversed Array is  : ", arr)

}
