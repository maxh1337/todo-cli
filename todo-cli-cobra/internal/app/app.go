package app

import (
	"todo-cli-cobra/cmd"
	"todo-cli-cobra/internal/storage"

	"github.com/spf13/cobra"
)

type App struct {
	Storage *storage.Storage
	RootCmd *cobra.Command
}

func NewApp() *App {
	app := &App{
		Storage: storage.LoadFromFile(),
	}

	rootCmd := &cobra.Command{Use: "todo"}
	rootCmd.AddCommand(cmd.NewCreateCmd(app.Storage),
		cmd.NewListCmd(app.Storage),
		cmd.NewDoneCmd(app.Storage),
		cmd.NewDeleteCmd(app.Storage),
	)
	app.RootCmd = rootCmd

	return app
}

func (a *App) Execute() error {
	return a.RootCmd.Execute()
}
