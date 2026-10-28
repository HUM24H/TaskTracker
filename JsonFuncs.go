package main

import (
	"encoding/json"
	"os"
)

func CheckJSONFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateJSONFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func CheckJSON(jsonString string) bool {
	var js map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &js)
	if err != nil {
		return false
	}
	return true
}

func WriteJSONToFile(filePath string, jsonData map[string]interface{}) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSONFromFile(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var js map[string]interface{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&js)
	if err != nil {
		return nil, err
	}

	return js, nil
}
