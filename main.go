package main

import (
	"database/sql"
	"log"

	"github.com/fadhilradh/simplybank/api"
	db "github.com/fadhilradh/simplybank/db/sqlc"
	"github.com/fadhilradh/simplybank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration")
	}
	conn, err := sql.Open(config.DBDriver, config.DSN)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
