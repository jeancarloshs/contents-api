package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("Connecting to the database...")
	dsn := "userdb:1234@tcp(localhost:3306)/vod_db"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
