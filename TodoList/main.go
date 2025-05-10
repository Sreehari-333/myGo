package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var task []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to TodoList")

	for {
		fmt.Println("\nSelect the Action")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Task")
		fmt.Println("3. Edit Task")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid choice, please enter a number")
			continue
		}

		switch choice {
		case 1:
			addTask(scanner)
		case 2:
			listTask()
		case 3:
			editTask(scanner)
		case 4:
			removeTask(scanner)
		case 5:
			fmt.Println("Exit from App")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addTask(scanner *bufio.Scanner) {

	fmt.Println("Enter the Description")
	scanner.Scan()

	desc := scanner.Text()
	task = append(task, desc)

	fmt.Println("Task Added")
}

func listTask() {

	if len(task) == 0 {
		fmt.Println("No task")
	} else {
		fmt.Println("Task List")
		for i, t := range task {
			fmt.Printf("%d. %s \n", i+1, t)
		}
	}

}

func editTask(scanner *bufio.Scanner) {
	fmt.Println("Enter the number of the task to be edit ")
	scanner.Scan()
	numEdit, err := strconv.Atoi(scanner.Text())
	if err != nil || numEdit <= 0 || numEdit > len(task) {
		fmt.Println("Invalid task number")
		return
	}

	fmt.Println("Enter the edited task ")
	scanner.Scan()
	editedTask := scanner.Text()

	task[numEdit-1] = editedTask
	fmt.Println("Task edited successfully")
}

func removeTask(scanner *bufio.Scanner) {

	fmt.Println("Enter the task to remove")
	scanner.Scan()

	removedTask, err := strconv.Atoi(scanner.Text())

	if err != nil || removedTask <= 0 || removedTask > len(task) {

		fmt.Println("invalid Task")
	} else {
		task = append(task[:removedTask-1], task[removedTask:]...)
		fmt.Println("Task removed")
	}

}
