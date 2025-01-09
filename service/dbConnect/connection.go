package dbConnect

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	// Build connection string using environment variables
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"

	// Open database connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Optional: Log the port for debugging
	log.Printf("App running on port: %s", port)
	return db
}
