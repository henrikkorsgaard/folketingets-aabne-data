package ftoda

import (
	"os"
	"sync"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

)

var (
	db *Repository
	dbOnce 	sync.Once
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	dbOnce.Do(func(){
		dbg, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DATABASE_PATH")),  &gorm.Config{})
		if err != nil {
			// we want to panic here because there is zero chance of recovering from a faulty db config/setup
			panic(err)

		}

		db = &Repository{dbg}
	})
	
	return db
}