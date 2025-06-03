package main

import (
	"todo-cli-cobra/internal/app"
)

func main() {
	a := app.NewApp()
	a.Execute()
}
