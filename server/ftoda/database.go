package ftoda

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dbRepository struct {
	db *gorm.DB
}

func newDBRepository(dbfile string) *dbRepository {
	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Sag{})
	return &dbRepository{
		db: db,
	}
}

func (db *dbRepository) insertBulk(sager []Sag) int64 {
	result := db.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&sager)
	if result.Error != nil {
		panic("error from insertBulk: " + result.Error.Error())
	}
	return result.RowsAffected
}

func (db *dbRepository) getRowCount(table string) int64 {
	var count int64
	db.db.Table(table).Count(&count)
	fmt.Println(count)
	return count
}
