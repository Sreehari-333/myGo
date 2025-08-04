package main

import "fmt"

func bubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {

	var n int

	fmt.Println("Enter the number of elements : ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the elements")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	bubbleSort(arr)
}
