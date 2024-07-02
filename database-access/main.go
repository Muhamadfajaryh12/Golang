package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
    // Load environment variables from .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get environment variables
    dbUser := os.Getenv("DBUSER")
    dbPass := os.Getenv("DBPASS")

    // Check if DBUSER and DBPASS are not empty
    if dbUser == "" || dbPass == "" {
        log.Fatal("DBUSER or DBPASS environment variable is not set")
    }

    // MySQL connection configuration
    dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/recordings", dbUser, dbPass)

    // Open database connection
    var errDB error
    db, errDB = sql.Open("mysql", dataSourceName)
    if errDB != nil {
        log.Fatal(errDB)
    }

    // Ping database to check connection
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

    fmt.Println("Connected to MySQL database!")
}
