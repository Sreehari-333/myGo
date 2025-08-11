package main

import (
	"container/heap"
	"fmt"
)

// Min Heap

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }           // len
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }      // smaller root
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] } // Swap

func (h *MinHeap) Push(x any) {

	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// Max Heap

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	minH := MinHeap{3, 2, 4, 1}
	heap.Init(&minH)
	heap.Push(&minH, 10)
	heap.Push(&minH, 5)
	fmt.Println(minH)
	fmt.Println("ExtractMin:", heap.Pop(&minH))

	maxH := &MaxHeap{}
	heap.Init(maxH)
	heap.Push(maxH, 10)
	heap.Push(maxH, 5)
	heap.Push(maxH, 20)
}
