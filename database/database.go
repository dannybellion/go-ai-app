package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./memory.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create tables
	CreateTables()
}

func CreateTables() {
	query := `CREATE TABLE IF NOT EXISTS chat_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		prompt TEXT,
		response TEXT
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
}