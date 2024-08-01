package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dataSources struct {
	DB *sqlx.DB
}

// InitDS establishes connections to fields in dataSources
func initDS() (*dataSources, error) {
	log.Printf("Initializing data sources\n")

	pgHost := getEnv("PG_HOST", "localhost")
	pgPort := getEnv("PG_PORT", "5432")
	pgUser := getEnv("PG_USER", "postgres")
	pgPassword := getEnv("PG_PASSWORD", "0000")
	pgDB := getEnv("PG_DB", "postgres")
	pgSSL := getEnv("PG_SSL", "disable")

	pgConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDB, pgSSL)

	log.Printf("Connecting to PostgreSQL\n")
	db, err := sqlx.Open("postgres", pgConnString)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working with retry logic
	for i := 0; i < 5; i++ {
		if err := db.Ping(); err != nil {
			log.Printf("Error connecting to db: %v, retrying in 2 seconds...", err)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db after retries: %w", err)
	}

	log.Printf("Successfully connected to PostgreSQL\n")
	return &dataSources{
		DB: db,
	}, nil
}

// close to be used in graceful server shutdown
func (d *dataSources) close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing PostgreSQL: %w", err)
	}

	return nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
