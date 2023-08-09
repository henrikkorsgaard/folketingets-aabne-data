package resolver

import (
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
	dbOnce.Do(func(){
		db, err := sql.Open("sqlite3", os.Getenv("SQLITE_DATABASE_PATH"))
		if err != nil {
			// we want to panic here because there is zero chance of recovering from a faulty db config/setup
			panic(err)

		}

		sqlitedb = &sqlite{db}
	})
	
	return sqlitedb
}