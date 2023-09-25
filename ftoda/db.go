package ftoda

import (
	"os"
	"sync"
	"gorm.io/gorm/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	repo   *Repository
	dbOnce sync.Once
	withlog bool = true
)

type Repository struct {
	db *gorm.DB
}

func newRepository() *Repository {
	// I'm using the singleton pattern here because we only do reads
	// and GORM chain methods get a new DB object anyways
	dbOnce.Do(func() {
		config := gorm.Config{}
		if withlog {
			config = gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			}
		}

		dbg, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DATABASE_PATH")), &config)
		if err != nil {
			// we want to panic here because there is zero chance of recovering from a faulty db config/setup
			panic(err)

		}

		repo = &Repository{dbg}
	})

	return repo
}
