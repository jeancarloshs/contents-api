package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
}

func conn() config {
	return config{
		dbUser:     "userdb",
		dbPassword: "1234",
		dbHost:     "localhost",
		dbPort:     "3306",
		dbName:     "vod_db",
	}
}

func ConnectDB() (*sql.DB, error) {
	config := conn()

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.dbUser, config.dbPassword, config.dbHost, config.dbPort, config.dbName)
	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Printf("Erro ao abrir a conexão: %v\n", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("Erro ao conectar ao banco: %v\n", err)
		return nil, err
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return db, nil
}
