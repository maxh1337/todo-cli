package cmd

import (
	"fmt"
	"todo-cli-cobra/internal/storage"
	"todo-cli-cobra/utils"

	"github.com/spf13/cobra"
)

func NewListCmd(storage *storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Вывести задачи",
		Run: func(cmd *cobra.Command, args []string) {
			utils.TitlePrint.Printf("Showing all Tasks: \n")
			fmt.Printf("\n")
			for i := 0; i < len(storage.Tasks); i++ {
				todo := (storage.Tasks)[i]
				var status string
				if todo.Done {
					status = "✅"
					utils.GreenPrint.Printf("%s ID: [%d] Title: %s\n", status, todo.ID, todo.Text)
				} else {
					status = "❌"
					utils.RedPrint.Printf("%s ID: [%d] Title: %s\n", status, todo.ID, todo.Text)
				}

			}
			fmt.Printf("\n")
		},
	}
}
