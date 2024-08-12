package psql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func init() {
	// Get the database URL from the environment variable
	dbURL := os.Getenv("DATABASE_URL")

	// Open a connection to the database
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set the global DB variable
	DB = db
}
