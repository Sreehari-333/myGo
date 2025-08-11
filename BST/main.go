package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

// Insert

func insert(root *Node, value int) *Node {

	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}
	return root
}

// InOrder

func inorder(root *Node) {
	if root != nil {
		inorder(root.left)
		fmt.Print(root.value, " ")
		inorder(root.right)
	}
}

// PreOrder

func preOrder(root *Node) {

	if root != nil {
		fmt.Println(root.value, " ")
		preOrder(root.left)
		preOrder(root.right)
	}

}

// PostOrder

func PostOrder(root *Node) {

	if root != nil {
		PostOrder(root.left)
		PostOrder(root.right)
		fmt.Println(root.value, " ")
	}
}

// Search in BST

func search(root *Node, value int) bool {

	if root == nil {
		return false
	}

	if value == root.value {
		return true
	} else if value < root.value {
		return search(root.left, value)
	} else {
		return search(root.right, value)
	}
}

// Find Min

func findMin(node *Node) *Node {

	current := node

	for current != nil && current.left != nil {
		current = current.left
	}
	return current
}

// Delete

func deleteNode(root *Node, value int) *Node {

	if root == nil {
		return root
	}

	if value < root.value {
		return deleteNode(root.left, value)
	} else if value > root.value {
		return deleteNode(root.right, value)
	} else {

		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		temp := findMin(root.right)
		root.value = temp.value
		root.right = deleteNode(root.right, temp.value)
	}
	return root
}

func main() {

	var root *Node

	root = insert(root, 10)
	root = insert(root, 20)
	root = insert(root, 5)

	// preOrder(root)
	// PostOrder(root)
	// deleteNode(root, 10)
	// inorder(root)

}
