// package main

// import "fmt"

// type Node struct {
// 	value int
// 	left  *Node
// 	right *Node
// }

// func insert(root *Node, value int) *Node {

// 	if root == nil {
// 		return &Node{value: value}
// 	}

// 	if value < root.value {
// 		root.left = insert(root.left, value)
// 	} else {
// 		root.right = insert(root.right, value)
// 	}

// 	return root
// }

// func inOrder(root *Node) {
// 	if root != nil {
// 		inOrder(root.left)
// 		fmt.Println(root.value, " ")
// 		inOrder(root.right)
// 	}
// }

// func preOrder(root *Node) {
// 	if root != nil {
// 		fmt.Println(root.value, " ")
// 		preOrder(root.left)
// 		preOrder(root.right)
// 	}
// }

// func postOrder(root *Node) {
// 	if root != nil {
// 		postOrder(root.left)
// 		postOrder(root.right)
// 		fmt.Println(root.value, " ")
// 	}
// }

// func search(root *Node, value int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if value == root.value {
// 		return true
// 	} else if value < root.value {
// 		return search(root.left, value)
// 	} else {
// 		return search(root.right, value)
// 	}
// }

// func findMin(node *Node) *Node {
// 	current := node

// 	for current != nil && current.left != nil {
// 		current = current.left
// 	}
// 	return current
// }

// func deleteNode(root *Node, value int) *Node {
// 	if root == nil {
// 		return root
// 	}

// 	if value < root.value {
// 		return deleteNode(root.left, value)
// 	} else if value > root.value {
// 		return deleteNode(root.right, value)
// 	} else {
// 		if root.left == nil {
// 			return root.right
// 		} else if root.right == nil {
// 			return root.left
// 		}

// 		temp := findMin(root.right)
// 		root.value = temp.value
// 		root.right = deleteNode(root.right, temp.value)
// 	}
// 	return root
// }
// func main() {

// 	var root *Node

// 	root = insert(root, 10)
// 	root = insert(root, 20)
// 	root = insert(root, 15)

// 	// preOrder(root)
// 	inOrder(root)
// 	// postOrder(root)

// 	// fmt.Println(search(root, 10))
// 	fmt.Println(deleteNode(root, 15)) // fmt.Println(root.right)
// }

package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func main() {
	minH := MinHeap{3, 4, 2, 5, 1}
	heap.Init(&minH)
	heap.Push(&minH, 10)
	fmt.Println(minH)
	fmt.Println(heap.Pop(&minH))
}
