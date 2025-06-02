package main

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-cli/commands"
)

func main() {
	file, err := os.ReadFile("./todos.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var storage commands.Storage
	err = json.Unmarshal(file, &storage)
	if err != nil {
		fmt.Println(err)
		return
	}

	commands.ExecuteFlags(&storage, os.Args[1:])

	// Перезаписываем файл, с измененными данными каждый раз после завершения работы
	newJsonFile, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile("./todos.json", newJsonFile, 0644)
}
