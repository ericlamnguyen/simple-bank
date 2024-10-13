package main

import (
	"database/sql"
	"log"

	"github.com/ericlamnguyen/simple-bank/api"
	db "github.com/ericlamnguyen/simple-bank/db/sqlc"
	"github.com/ericlamnguyen/simple-bank/util"

	_ "github.com/lib/pq"
)

func main() {
	// Read config file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Create a connection pool to the db
	connPool, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	defer connPool.Close()

	// Create store object to interact with the db
	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
