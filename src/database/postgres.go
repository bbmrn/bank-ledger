package database

import (
	"database/sql"
	"fmt"
	"log"

	"bank-ledger/util"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

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
