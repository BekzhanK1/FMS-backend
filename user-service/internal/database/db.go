package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	config "user-service/internal/config"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitStorage() (*sql.DB, error){
	dbconfig := config.Envs

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.DBHost,
		dbconfig.DBPort,
		dbconfig.DBUser,
		dbconfig.DBPassword,
		dbconfig.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to opend database: %v", err)
	}

	return db, nil
}

func Connect() (*sql.DB, error) {
	db, err := InitStorage()
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %v", err)
	}

	log.Println("Database connected successfully.")

	return db, nil
}