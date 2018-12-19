package config

import (
	"database/sql"
	"log"

	_ "github.com/bmizerany/pq"
)

// NewPGConnection create new postgres database connection
// it return postgres connection and cleanup postgres connection
func NewPGConnection(dataSource string) (*sql.DB, func()) {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db, func() {
		err := db.Close()
		if err != nil {
			log.Println("Failed to close DB by error", err)
		}
	}
}
