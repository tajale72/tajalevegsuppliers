package aiven

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDBAivenPostgres() (*sql.DB, error) {
	// Load sensitive data from environment variables
	username := os.Getenv("AIVEN_PG_USERNAME")
	password := os.Getenv("AIVEN_PG_PASSWORD")
	host := os.Getenv("AIVEN_PG_HOST")
	port := os.Getenv("AIVEN_PG_PORT")
	dbName := os.Getenv("AIVEN_PG_DB")
	sslRootCert := os.Getenv("AIVEN_PG_SSLROOTCERT")

	if username == "" || password == "" || host == "" || port == "" || dbName == "" || sslRootCert == "" {
		return nil, fmt.Errorf("one or more environment variables for the DB connection are not set")
	}

	// Construct the connection string
	serviceURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=verify-ca&sslrootcert=%s",
		username, password, host, port, dbName, sslRootCert,
	)

	// Open a connection to the Aiven Postgres database
	db, err := sql.Open("postgres", serviceURI)
	if err != nil {
		log.Printf("Error opening Aiven Postgres connection: %v", err)
		return nil, err
	}

	// Ping the database to ensure the connection is valid
	if err := db.Ping(); err != nil {
		log.Printf("Error pinging Aiven Postgres database: %v", err)
		return nil, err
	}

	log.Println("Connected to Aiven Postgres successfully")
	return db, nil
}

