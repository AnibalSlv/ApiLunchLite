/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"apiLunchLite/cmd"
	"apiLunchLite/internal/database"
	"apiLunchLite/internal/services"
	"apiLunchLite/internal/services/manager"
	"apiLunchLite/models"
	"log"

	"database/sql"
)

func main() {
	conn, err := sql.Open("sqlite", "DbAPL.db")
	if err != nil {
		log.Fatal("Error Open Db:", err)
	}
	defer conn.Close()

	repo := &database.SQLite{DbConn: conn}

	// Se crea un crotato (interfaz)
	var db models.Db = repo

	db.InitDb()

	miLogger := &services.Logger{
		FolderLogs: "internal/logs",
	}

	ApiManager := &manager.ApiManager{
		Logger: miLogger,
		Db:     db,
	}

	cmd.Execute(ApiManager)
}
