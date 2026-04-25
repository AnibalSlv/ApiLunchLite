/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"apiLunchLite/cmd"
	"apiLunchLite/internal/database"
	"log"

	"database/sql"
)

func main() {
	db, err := sql.Open("sqlite", "DbAPL.db")
	if err != nil {
		log.Fatal("Error Open Db:", err)
	}
	defer db.Close()

	database.InitDb()

	cmd.SetDb(db)

	cmd.Execute()
}
