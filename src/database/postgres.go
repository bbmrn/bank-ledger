package database

import (
	"database/sql"
	"fmt"
	"log"

	"bank-ledger/util"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

// InitPostgres initializes a connection to PostgreSQL database using connection parameters
// from environment variables or default values. It establishes and verifies the connection
// by attempting to ping the database.
//
// The function uses the POSTGRES_URL environment variable for connection string, falling back
// to default credentials if not set. It sets up the global PostgresDB variable for database
// operations throughout the application.
//
// The function will terminate the program with a fatal error if:
// - Connection cannot be established with the database
// - Database ping fails
func InitPostgres() {
	connStr := util.GetEnv("POSTGRES_URL", "user=youruser dbname=yourdb sslmode=disable password=yourpassword")
	var err error
	PostgresDB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}

	err = PostgresDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging PostgreSQL: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}

func BeginTransaction() (*sql.Tx, error) {
	return PostgresDB.Begin()
}
