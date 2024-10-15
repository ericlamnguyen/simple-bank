package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ericlamnguyen/simple-bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testStore Store

// TestMain is the main entry point for the tests in this package.
func TestMain(m *testing.M) {
	// Load config file
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Setup code (run before any tests)
	connPool, err := sql.Open(config.DBDriver, config.DBSource)
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
