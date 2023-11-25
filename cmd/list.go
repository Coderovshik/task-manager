package cmd

import (
	"fmt"

	"github.com/Coderovshik/task-manager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks on your toDo list",
	Long:  "list all tasks on your toDo list",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have the following tasks:")
		taskList := db.GetTaskList()
		sortedTasks := make([][]string, len(taskList))
		for k, v := range taskList {
			// fmt.Printf("%d. %s\n", k, v[1])
			sortedTasks[k-1] = v
		}

		for i, v := range sortedTasks {
			fmt.Printf("%d. %s\n", i+1, v[1])
		}
	},
}
