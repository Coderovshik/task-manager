package main

import (
	"github.com/Coderovshik/task-manager/cmd"
	"github.com/Coderovshik/task-manager/db"
)

func main() {
	db.Open()
	defer db.Close()

	cmd.Execute()
}
