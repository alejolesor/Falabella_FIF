package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func NewConnectionDB() *sql.DB {
	// Open up our database connection.
	conn := fmt.Sprintf("bender:pass1234@tcp(%s:3306)/beer?charset=utf8mb4&parseTime=True&loc=Local&timeout=1m", dbHost())
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed ping: %+v", err)
	}

	return db
}

func dbHost() string {
	if dbHost, exist := os.LookupEnv("DB_HOST"); exist {
		return dbHost
	}

	return "localhost"
}
