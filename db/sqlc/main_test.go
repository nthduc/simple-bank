package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_"github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:021220@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// connection
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
