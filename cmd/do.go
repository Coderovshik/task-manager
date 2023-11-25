package cmd

import (
	"fmt"
	"strconv"

	"github.com/Coderovshik/task-manager/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do [task list]",
	Short: "Mark task as completed",
	Long:  "Mark task as completed",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		if v, ok := db.GetTaskList()[id]; ok {
			task := v[1]
			db.RemoveTask(v[0])
			fmt.Printf("You have completed the %q task.\n", task)
		}
	},
}
