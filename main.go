package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const fileName = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [add|list|update|delete|mark-in-progress|mark-done]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	default:
		fmt.Println("Unknown command")
	}
}

func addTask(desc string) {
	tasks := loadTasks()

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
}

func loadTasks() []Task {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func listTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range tasks {
		fmt.Printf("[%d] %s - %s\n", t.ID, t.Description, t.Status)
	}
}