package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
const tasksFile = "tasks.json" // File for persisting tasks 

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func loadTasks() ([]Task, error) {
	return nil, nil
}

func saveTasks(tasks []Task) error {
	// write tasks to file
	data, err := json.Marshal(tasks)
	if err != nil {	
		return err
	}
	return ioutil.WriteFile(tasksFile, data, 0644)	
}

func listTasks(tasks []Task) {}

func addTask(tasks []Task, description string) []Task {
	tasks = append(tasks, Task{Description: description, Completed: false})
	return tasks
}

func completeTask(tasks []Task, index int) ([]Task, error) {
	return nil, nil
}

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Usage: taskcli [add|list|done] [task description|task number]")
		return
	}

	var action string
	var taskDescription string
	action = os.Args[1]

	// load tasks from file, as this will be used for all actions
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	switch action {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskcli add [task description]")
			return
		}
		taskDescription = os.Args[2]
		tasks = addTask(tasks, taskDescription)
		// Save the updated tasks
		if err := saveTasks(tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			return
		}
		fmt.Println("Task added:", taskDescription)
	case "list":
	case "done":
	default:
		fmt.Println("Unknown action:", action)
		return
	}
}