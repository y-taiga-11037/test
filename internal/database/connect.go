package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	// Read the env file
	err := godotenv.Load("/home/y-taiga/mdtd_bootcamp/env/dev.env")
	if err != nil {
		log.Fatal(err)
	}

	// Login to mysql
	db, err := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
