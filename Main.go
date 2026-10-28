package main

import (
	"fmt"
	"os"
)

var jsonFile string = "tasks.json"

func main() {
	if CheckJSONFileExists(jsonFile) {
		fmt.Println("JSON file exists.")
	} else {
		fmt.Println("JSON file does not exist.")
		if err := CreateJSONFile(jsonFile); err != nil {
			fmt.Println("Error creating JSON file:", err)
			return
		} else {
			fmt.Println("JSON file created successfully.")
		}
	}

	checkArgs := len(os.Args)
	if checkArgs < 3 {
		fmt.Println("Usage: go run Main.go <mode> <task_name>")
		fmt.Println("Modes: add, update, delete")
		return
	}

	mode := os.Args[1]
	taskNameOrID := os.Args[2]
	taskUpdateVal := ""
	listMode := ""

	if checkArgs > 3 && mode == "update" {
		taskUpdateVal = os.Args[3]
	} else if checkArgs > 2 && mode == "list" {
		listMode = os.Args[2]
	} else if checkArgs > 3 && mode == "update-status" {
		taskUpdateVal = os.Args[3]
	}

	if taskNameOrID == "" {
		fmt.Println("Task name or ID is required.")
		return
	}

	switch mode {
	case "add":
		// Add
		tasks, err := ReadJSONFromFile(jsonFile)
		if err != nil {
			if tasks != nil {
				fmt.Println("Error reading JSON file:", err)
				return
			} else {
				tasks = make(map[string]interface{})
			}
		}

		taskID := GetNextFreeTaskID(tasks)
		newTask := map[string]interface{}{
			"id":     taskID,
			"name":   taskNameOrID,
			"status": "pending",
		}

		tasks[taskID] = newTask

		if err := WriteJSONToFile(jsonFile, tasks); err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return
		}
		fmt.Println("Task added successfully.")
	case "update":
		// Update
		tasks, err := ReadJSONFromFile(jsonFile)
		if err != nil {
			fmt.Println("Error reading JSON file:", err)
			return
		}

		task, exists := tasks[taskNameOrID]
		if !exists {
			fmt.Println("Task with ID", taskNameOrID, "does not exist.")
			return
		}

		if taskMap, ok := task.(map[string]interface{}); ok {
			taskMap["name"] = taskUpdateVal
			tasks[taskNameOrID] = taskMap
		} else {
			fmt.Println("Invalid task format.")
			return
		}

		if err := WriteJSONToFile(jsonFile, tasks); err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return
		}
		fmt.Println("Task updated successfully.")
	case "delete":
		// Delete
		tasks, err := ReadJSONFromFile(jsonFile)
		if err != nil {
			fmt.Println("Error reading JSON file:", err)
			return
		}

		if _, exists := tasks[taskNameOrID]; !exists {
			fmt.Println("Task with ID", taskNameOrID, "does not exist.")
			return
		}

		delete(tasks, taskNameOrID)

		if err := WriteJSONToFile(jsonFile, tasks); err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return
		}
		fmt.Println("Task deleted successfully.")
	case "update-status":
		// Update task status
		if taskUpdateVal != "done" && taskUpdateVal != "pending" && taskUpdateVal != "in-progress" {
			fmt.Println("Invalid status value. Please use 'done', 'pending', or 'in-progress'.")
			return
		}

		tasks, err := ReadJSONFromFile(jsonFile)
		if err != nil {
			fmt.Println("Error reading JSON file:", err)
			return
		}

		task, exists := tasks[taskNameOrID]
		if !exists {
			fmt.Println("Task with ID", taskNameOrID, "does not exist.")
			return
		}

		if taskMap, ok := task.(map[string]interface{}); ok {
			taskMap["status"] = taskUpdateVal
			tasks[taskNameOrID] = taskMap
		} else {
			fmt.Println("Invalid task format.")
			return
		}

		if err := WriteJSONToFile(jsonFile, tasks); err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return
		}
		fmt.Println("Task status updated successfully.")
	case "list":
		// List tasks
		if listMode != "done" && listMode != "todo" && listMode != "in-progress" && listMode != "all" {
			fmt.Println("Invalid list mode. Please use 'done', 'todo', 'in-progress', or 'all'.")
			return
		}

		tasks, err := ReadJSONFromFile(jsonFile)
		if err != nil {
			fmt.Println("Error reading JSON file:", err)
			return
		}

		var checkStatus string

		switch listMode {
		case "done":
			checkStatus = "done"
		case "todo":
			checkStatus = "pending"
		case "in-progress":
			checkStatus = "in-progress"
		case "all":
			checkStatus = ""
		default:
			checkStatus = ""
		}

		for _, task := range tasks {
			if taskMap, ok := task.(map[string]interface{}); ok {
				if status, ok := taskMap["status"].(string); ok && (status == checkStatus || checkStatus == "") {
					fmt.Printf("ID: %s, Name: %s, Status: %s\n", taskMap["id"], taskMap["name"], taskMap["status"])
				}
			}
		}
	default:
		fmt.Println("Invalid mode. Please use 'add', 'update', or 'delete'.")
	}
}
