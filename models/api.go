package models

import "database/sql"

type ApiConfig struct {
	DbConn     *sql.DB
	Id         int
	Name       string
	Host       string
	Port       int
	PathFolder string
	State      string
	Pid        int
}
