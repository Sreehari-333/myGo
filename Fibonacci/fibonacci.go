package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a task in the ToDo list
type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks []Task
var nextID = 1

// Show all tasks
func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}
	for _, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Description, status)
	}
}

// Add a new task
func addTask(description string) {
	task := Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added!")
}

// Mark a task as completed
func completeTask(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

// Delete a task
func deleteTask(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTodo List Command-Line Tool")
		fmt.Println("1. Show Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter command: ")

		scanner.Scan()
		command := scanner.Text()

		switch strings.ToLower(command) {
		case "1":
			showTasks()
		case "2":
			fmt.Print("Enter task description: ")
			scanner.Scan()
			description := scanner.Text()
			addTask(description)
		case "3":
			fmt.Print("Enter task ID to mark as completed: ")
			scanner.Scan()
			var taskID int
			_, err := fmt.Sscanf(scanner.Text(), "%d", &taskID)
			if err != nil {
				fmt.Println("Invalid input")
				break
			}
			completeTask(taskID)
		case "4":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			var taskID int
			_, err := fmt.Sscanf(scanner.Text(), "%d", &taskID)
			if err != nil {
				fmt.Println("Invalid input")
				break
			}
			deleteTask(taskID)
		case "5":
			fmt.Println("Exiting Todo List Tool.")
			return
		default:
			fmt.Println("Invalid command! Please try again.")
		}
	}
}
