package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection ping failed: ", err)
	}

	log.Println("Successfully connected to PostgreSQL using PGX")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to run database migration: ", err)
	}

	log.Println("Database migration applied successfully")
}
