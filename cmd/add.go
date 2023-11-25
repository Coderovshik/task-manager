package cmd

import (
	"fmt"
	"strings"

	"github.com/Coderovshik/task-manager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add task to your task list",
	Long:  "Add task to your task list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		db.AddTask(task)
		fmt.Printf("Added %q to your task list.\n", task)
	},
}
