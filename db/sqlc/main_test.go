package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/otaviomart1ns/finsys/config"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer conn.Close()

	testQueries = New(conn)
	os.Exit(m.Run())
}
