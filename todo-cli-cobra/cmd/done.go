package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"todo-cli-cobra/internal/storage"
	"todo-cli-cobra/utils"

	"github.com/spf13/cobra"
)

func NewDoneCmd(store *storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use:   "done",
		Short: "Изменить статус задачи",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				utils.ErrorPrint.Printf("Please provide task id to toggle. Example: done 1\n")
				return
			}
			text := args[0]
			taskID, err := strconv.Atoi(text)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
			index := slices.IndexFunc(store.Tasks, func(task storage.Task) bool {
				return task.ID == taskID
			})

			if index == -1 {
				utils.ErrorPrint.Printf("Task not found, provide correct task ID \n")
				return
			}

			if store.Tasks[index].Done {
				utils.ErrorPrint.Printf("This task is already DONE, try completing another task or delete this\n")
				return
			}

			store.Tasks[index].Done = true

			storage.SaveResults(store)
			utils.SuccessPrint.Printf("Task #%v delete\n", taskID)
		},
	}
}
