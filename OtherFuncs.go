package main

import "fmt"

func GetNextFreeTaskID(tasks map[string]interface{}) string {
	uniqueIDs := GetUniqueIDs(tasks)

	for i := 1; ; i++ {
		id := fmt.Sprintf("%d", i)
		if !uniqueIDs[id] {
			return id
		}
	}
}

func GetUniqueIDs(tasks map[string]interface{}) map[string]bool {
	uniqueIDs := make(map[string]bool)
	for _, task := range tasks {
		if taskMap, ok := task.(map[string]interface{}); ok {
			if id, ok := taskMap["id"].(string); ok {
				uniqueIDs[id] = true
			}
		}
	}
	return uniqueIDs
}
