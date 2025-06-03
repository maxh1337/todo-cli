package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Storage struct {
	NextID int
	Tasks  []Task
}

type Task struct {
	ID   int
	Text string
	Done bool
}

var FilePath string = filepath.Join("internal", "storage", "todos.json")

func LoadFromFile() *Storage {
	var storage Storage
	file, err := os.ReadFile(FilePath)
	if err != nil {
		fmt.Printf("while loading file, error occurred - %s \n", err)
	}

	json.Unmarshal(file, &storage)

	return &storage
}

func IncrementStorageId(storage *Storage) error {
	storage.NextID += 1
	return nil
}

func SaveResults(storage *Storage) error {
	newJsonFile, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// path := filepath.Join("internal", "storage", "todos.json")
	err = os.WriteFile(FilePath, newJsonFile, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
