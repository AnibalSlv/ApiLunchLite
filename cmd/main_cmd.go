package cmd

import (
	"database/sql"

	"apiLunchLite/models"

	_ "modernc.org/sqlite"
)

var Api models.ApiConfig

func SetDb(db *sql.DB) {
	Api.DbConn = db
}
