package database

import (
	"database/sql"
	"os"

	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	logging "github.com/sirupsen/logrus"
)

var DB *sql.DB

func Connect() *sql.DB {

	/* TODO:
	drone test does not recognize file path
	The log file and the environment setting file are read by the CLI
	Fix it when the CLI is installed


	// Read the env file
	/*	err := godotenv.Load("/home/y-taiga/mdtd_bootcamp/env/dev.env")
		if err != nil {
			logging.Error("Unable to read the database configuration file")
		}
		logging.Info("Read the database configuration file") */

	// Login to mysql
	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		logging.Error("Unable to connect to the database")
	}
	logging.Info("Connect to the database")

	return db
}
