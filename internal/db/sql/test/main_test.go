package db_test

import (
	"database/sql"
	db "highscores-core/internal/db/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/highscores_core?sslmode=disable"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}
	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
