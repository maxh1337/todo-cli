package cmd

import (
	"todo-cli-cobra/internal/storage"
	"todo-cli-cobra/utils"

	"github.com/spf13/cobra"
)

func NewCreateCmd(store *storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Создать задачу",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				utils.ErrorPrint.Printf("Please provide title for task create 'title'\n")
				return
			}
			text := args[0]

			if text == "" {
				utils.ErrorPrint.Printf("The string is empty, please provide correct title\n")
				return
			}

			oneNewTask := storage.Task{
				ID:   store.NextID,
				Text: text,
				Done: false,
			}

			storage.IncrementStorageId(store)
			store.Tasks = append(store.Tasks, oneNewTask)
			storage.SaveResults(store)

			utils.SuccessPrint.Printf("Task '%v' created\n", text)
		},
	}
}
