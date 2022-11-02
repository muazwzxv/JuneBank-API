package adapter

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"
)

var database *sql.DB

func GetDatabaseAdapter() *sql.DB {
	if database == nil {
		if conn, err := connect(); err != nil {
			log.Fatalf("Shutdown: cannot connect to db")
		} else {
			database = conn
		}
	}

	return database
}

func connect() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	const maxAttempt = 10
	attempt := 0

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Printf("failed to connect to db: %v", err)
			attempt++
		} else {
			log.Println("database connected")
			return conn, nil
		}

		if attempt == maxAttempt {
			return nil, errors.New("timeout")
		}

		time.Sleep(2 * time.Second)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
