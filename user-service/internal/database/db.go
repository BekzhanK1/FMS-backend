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

func Connect() (*sql.DB, error) {
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
		return nil, fmt.Errorf("Unable to connect to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("Unable to connect to the database: %v", err)	
	}

	log.Println("Database connected successfully.")

	return DB, nil
}
