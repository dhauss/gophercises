package main

import (
	"fmt"
	"os"
)

const default_help = `task is a CLI for managing your TODOs.
Usage:
	task [command]
	
Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks
	
Use "task [command] --help" for more information about a command.`

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(default_help)
		os.Exit(0)
	}

	switch args[0] {
	case "add":
		fmt.Println(default_help)
	case "list":
		fmt.Println("listin'")
	case "do":
		fmt.Println("doin'")
	default:
		fmt.Println(default_help)
	}
}
