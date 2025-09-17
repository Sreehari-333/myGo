package main

import "fmt"

// Splits  the array and sort recurssively

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)

}

// Merge combines two sorted slices

func merge(left, right []int) []int {

	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {

	arr := []int{7, 5, 2, 1, 3, 6, 4}
	fmt.Println("Orginal Array : ", arr)

	sorted := mergeSort(arr)
	fmt.Println("Sorted array : ", sorted)

}
