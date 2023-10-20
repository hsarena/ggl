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
	dbConnection = "postgresql://root:secret@localhost:5432/ggl?sslmode=disable"
)


var testQueries *Queries


func TestMain (m *testing.M) {
	conn, err := sql.Open(dbDriver, dbConnection)
	if err != nil {
		log.Fatal("could not connect to db: ", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
