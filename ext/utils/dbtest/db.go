package db

import (
	"database/sql"
	"fmt"
	"go/build"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	_ "github.com/lib/pq" // driver for open postgres connection
	"github.com/pressly/goose"
)

const (
	migrationSource = "/src/github.com/hieuphq/califit-be/migration"
	seedSource      = "/src/github.com/hieuphq/califit-be/migration/test"
	dbHost          = "localhost"
	dbPort          = 5439
	dbUser          = "postgres"
	dbPassword      = "example"
	dbName          = "test"
)

// CreateTestDatabase will create a test-database and test-schema
func CreateTestDatabase(t *testing.T) (*sql.DB, string, func()) {
	testingHost := fmt.Sprintf("%s", dbHost)
	testingPort := fmt.Sprintf("%d", dbPort)
	if os.Getenv("POSTGRES_TESTING_HOST") != "" {
		testingHost = os.Getenv("POSTGRES_TESTING_HOST")
	}
	if os.Getenv("POSTGRES_TESTING_PORT") != "" {
		testingPort = os.Getenv("POSTGRES_TESTING_PORT")
	}

	connectionString := fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=disable", testingHost, testingPort, dbUser, dbPassword, dbName)
	db, dbErr := sql.Open("postgres", connectionString)
	if dbErr != nil {
		t.Log(connectionString)
		t.Logf("host env: %v, port: %v", os.Getenv("POSTGRES_TESTING_HOST"), os.Getenv("POSTGRES_TESTING_PORT"))
		t.Fatalf("Fail to create database. %s", dbErr.Error())
	}
	defer db.Close()

	rand.Seed(time.Now().UnixNano())
	schemaName := "test" + strconv.FormatInt(rand.Int63(), 10)

	_, err := db.Exec("CREATE SCHEMA " + schemaName)
	if err != nil {
		t.Fatalf("Fail to create schema. %s", err.Error())
	}

	connectionString = fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", testingHost, testingPort, dbUser, dbPassword, dbName, schemaName)
	db, dbErr = sql.Open("postgres", connectionString)
	if dbErr != nil {
		t.Fatalf("Fail to create database. %s", dbErr.Error())
	}

	return db, schemaName, func() {
		_, err := db.Exec("DROP SCHEMA " + schemaName + " CASCADE")
		if err != nil {
			t.Fatalf("Fail to drop database. %s", err.Error())
		}
	}
}

// MigrateTest -- migrate data with migration file for testing with pg lib
func MigrateTest(db *sql.DB) error {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	//  create tables
	if err := goose.Up(db, gopath+migrationSource); err != nil {
		return err
	}

	// seed test data
	if err := goose.Up(db, gopath+seedSource); err != nil {
		return err
	}
	return nil
}
