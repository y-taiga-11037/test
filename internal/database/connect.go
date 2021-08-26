package database

import (
	"database/sql"
	"os"

	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	"github.com/joho/godotenv"
	logging "github.com/sirupsen/logrus"
)

func Connect() *sql.DB {

	// Read the env file
	err := godotenv.Load("/home/y-taiga/mdtd_bootcamp/env/dev.env")
	if err != nil {
		logging.Error("Unable to read the database configuration file")
	}
	logging.Info("Read the database configuration file")

	// Login to mysql
	db, err := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		logging.Error("Unable to connect to the database")
	}
	logging.Info("Connect to the database")

	return db
}
