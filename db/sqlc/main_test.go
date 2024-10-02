package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testStore *Store

// TestMain is the main entry point for the tests in this package.
func TestMain(m *testing.M) {
	// Setup code (run before any tests)
	connPool, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	defer connPool.Close()

	testQueries = New(connPool)
	testStore = NewStore(connPool)

	// Run the tests
	exitCode := m.Run()

	// Teardown code (run after all tests)

	// Exit with the code returned from m.Run()
	os.Exit(exitCode)
}
