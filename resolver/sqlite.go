package resolver

import (
	"fmt"
	"os"
	"sync"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

)

var (
	sqlitedb *sqlite
	dbOnce 	sync.Once
)

type sqlite struct {
	db *sql.DB
}

func (s *sqlite) Close(){
	s.db.Close()
}

func newSqlite() *sqlite {
	fmt.Println("hu")
	fmt.Println(os.Getenv("SQLITE_TEST_DATABASE_PATH"))
	dbOnce.Do(func(){
		db, err := sql.Open("sqlite3", "../ingest/data/odatest.sqlite.db")
		if err != nil {
			// we want to panic here because there is zero chance of recovering from a faulty db config/setup
			panic(err)

		}

		sqlitedb = &sqlite{db}
	})
	
	return sqlitedb
}