package repository

import (
	"errors"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrDatabaseConnection = errors.New("error connecting to the database")
	ErrDatabaseUpdateSag  = errors.New("error updating sager in database")
)

type dbRepository struct {
	db *gorm.DB
}

func newDatabaseRepo(host string) *dbRepository {
	db, err := gorm.Open(sqlite.Open(host), &gorm.Config{})
	if err != nil {
		panic(errors.Join(ErrDatabaseConnection, err))
	}

	db.AutoMigrate(&ftoda.Sag{})
	return &dbRepository{
		db: db,
	}
}

func (repo *dbRepository) getSagById(id int) (sag ftoda.Sag, err error) {
	result := repo.db.First(&sag, id)
	//this will only return errors (from result.Error) if the error is not
	//a record not found error
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return sag, errors.Join(ErrRepoGettingSag, err)
	}

	return sag, err
}

func (repo *dbRepository) getSagerByType(sagtype int) (sager []ftoda.Sag, err error) {

	result := repo.db.Where("typeid = ?", sagtype).Find(&sager)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return sager, errors.Join(ErrRepoGettingSag, err)
	}

	return sager, err
}

func (repo *dbRepository) updateSager(sager []ftoda.Sag) (rows int64, err error) {
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&sager)
	if result.Error != nil {
		return rows, errors.Join(ErrDatabaseUpdateSag, result.Error)
	}
	return result.RowsAffected, err
}
