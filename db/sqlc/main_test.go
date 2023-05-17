package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/fadhilradh/simplybank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DSN)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
