// internal/database/db.go

package database

import (
	"database/sql"
	"fmt"
	"log"

	config "user-service/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dbconfig := config.Load()
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.DBHost,
		dbconfig.DBPort,
		dbconfig.DBUser,
		dbconfig.DBPassword,
		dbconfig.DBName,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Unable to reach the database: %v", err)
	}

	log.Println("Database connected successfully.")
}
