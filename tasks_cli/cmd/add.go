package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {

		task := strings.Join(args, " ")
		db.ErrCatch(db.AddTask(task))

		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
