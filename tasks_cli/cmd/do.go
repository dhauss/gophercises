package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Removed \"%s\" from your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
