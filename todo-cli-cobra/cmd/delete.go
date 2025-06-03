package cmd

import (
	"slices"
	"strconv"
	"todo-cli-cobra/internal/storage"
	"todo-cli-cobra/utils"

	"github.com/spf13/cobra"
)

func NewDeleteCmd(store *storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Удалить задачу",
		Run: func(cmd *cobra.Command, args []string) {
			text := args[0]
			taskID, err := strconv.Atoi(text)
			if err != nil {
				utils.ErrorPrint.Printf("Error: %s", err)
				return
			}

			store.Tasks = slices.DeleteFunc(store.Tasks, func(todo storage.Task) bool {
				return todo.ID == taskID
			})

			storage.SaveResults(store)
			utils.SuccessPrint.Printf("Task #%v delete\n", taskID)
		},
	}
}
