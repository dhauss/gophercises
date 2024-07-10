package main

import (
	"log"
	"task/cmd"
	"task/db"
)

func main() {
	dbName := "tasks.db"

	errCatch(db.Init(dbName))
	errCatch(cmd.RootCmd.Execute())
}

func errCatch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
