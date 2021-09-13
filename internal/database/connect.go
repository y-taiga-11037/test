package database

import (
	"database/sql"
	"os"

	logging "github.com/sirupsen/logrus"
)

var DB *sql.DB

func Connect() *sql.DB {

	// Login to mysql
	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))

	if err != nil {
		logging.Error("Unable to connect to the database")
		os.Exit(1)
	}
	logging.Info("Connect to the database")

	return db
}
