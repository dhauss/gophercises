package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ListTasks()
		db.ErrCatch(err)

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Invalid task number %d \n", id)
				continue
			}
			curTask := tasks[id-1]
			err := db.DeleteTask(curTask.Key)

			if err != nil {
				fmt.Printf("Failed to remove \"%d\" from tasks. Error: %s\n", id, err)
			} else {
				fmt.Printf("Removed \"%d\" from your task list.\n", id)
			}

			err = db.CompleteTask(curTask.Value)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
