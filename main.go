package main

import (
	"database/sql"
	"log"

	"github.com/ericlamnguyen/simple-bank/api"
	db "github.com/ericlamnguyen/simple-bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddr = "0.0.0.0:8080"
)

func main() {
	// Create a connection pool to the db
	connPool, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	defer connPool.Close()

	// Create store object to interact with the db
	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
