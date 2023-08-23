package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitializeDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/avenger")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
