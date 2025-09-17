package main

import "fmt"

func quikSort(arr []int, low, high int) {

	if low < high {
		pi := partition(arr, low, high)
		quikSort(arr, low, pi-1)
		quikSort(arr, pi+1, high)
	}

}

func partition(arr []int, low, high int) int {

	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {

	arr := []int{3, 1, 4, 2}
	fmt.Println("Orginal Array : ", arr)

	quikSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted Array : ", arr)

}
