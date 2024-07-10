package main

import (
	"task/cmd"
	"task/db"
)

func main() {
	dbName := "tasks.db"

	db.ErrCatch(db.Init(dbName))
	db.ErrCatch(cmd.RootCmd.Execute())
}
