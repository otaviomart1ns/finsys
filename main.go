package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/otaviomart1ns/finsys/api"
	"github.com/otaviomart1ns/finsys/config"
	db "github.com/otaviomart1ns/finsys/db/sqlc"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(env.ServerAdress)
	if err != nil {
		log.Fatalf("cannot start api: %v", err)
	}
}
