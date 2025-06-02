package commands

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
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

func ExecuteFlags(storage *Storage, args []string) {
	fs := flag.NewFlagSet("todo", flag.ContinueOnError)

	fs.Func("done", "Example: -done 2", func(s string) error {
		id, err := strconv.Atoi(s)

		if err != nil {
			return fmt.Errorf("invalid task id: %v", err)
		}

		err = DoneTask(id, storage)
		if err != nil {
			return fmt.Errorf("smthing went wrong: %v", err)
		}

		return nil
	})

	fs.Func("delete", "Example: -delete 2", func(s string) error {
		id, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("invalid task id: %v", err)
		}

		err = DeleteTask(id, storage)
		if err != nil {
			return fmt.Errorf("deleting error: %v", err)
		}

		return nil
	})

	fs.BoolFunc("list", "-list Show all current tasks", func(s string) error {
		err := ListAllTasks(storage)
		if err != nil {
			return fmt.Errorf("it seems you doesn't have any todos rn: %v", err)
		}

		return nil
	})

	fs.Func("create", "Example: -create 'make homework'", func(s string) error {
		err := CreateTask(s, storage)
		if err != nil {
			return fmt.Errorf("creation error: %v", err)
		}

		return nil
	})

	fs.Usage = func() {
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("   -list Show all current tasks")
		fmt.Println("   -create <task name>  Create a new task")
		fmt.Println("   -done <task number>  Task done")
		fmt.Println("   -delete <task number>  Task deleted")
		fmt.Println("")
	}

	fs.Parse(args)
}

func CreateTask(taskName string, storage *Storage) error {
	oneNewTask := Task{
		ID:   storage.NextID,
		Text: taskName,
		Done: false,
	}

	IncrementStorageId(storage)
	storage.Tasks = append(storage.Tasks, oneNewTask)

	fmt.Printf("Task '%v' created\n", taskName)
	return nil
}

func DeleteTask(taskId int, storage *Storage) error {
	storage.Tasks = slices.DeleteFunc(storage.Tasks, func(todo Task) bool {
		return todo.ID == taskId
	})

	fmt.Printf("Task #%d deleted\n", taskId)
	return nil
}

func DoneTask(taskId int, storage *Storage) error {
	taskIndex := slices.IndexFunc(storage.Tasks, func(task Task) bool {
		return task.ID == taskId
	})

	(storage.Tasks)[taskIndex].Done = true

	fmt.Printf("Task slice index: %d\n", taskIndex)
	fmt.Printf("Task #%d marked as done\n", taskId)
	return nil
}

func ListAllTasks(storage *Storage) error {
	fmt.Printf("Showing all Tasks: \n")
	fmt.Printf("\n")
	for i := 0; i < len(storage.Tasks); i++ {
		todo := (storage.Tasks)[i]
		var status string
		if todo.Done {
			status = "✅"
		} else {
			status = "❌"
		}
		fmt.Printf("%s ID: [%d] Title: %s\n", status, todo.ID, todo.Text)
	}
	fmt.Printf("\n")
	return nil
}

func IncrementStorageId(storage *Storage) error {
	storage.NextID += 1
	return nil
}
