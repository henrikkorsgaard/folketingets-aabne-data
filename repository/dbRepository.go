package repository

import (
	"errors"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrDatabaseConnection          = errors.New("error connecting to the database")
	ErrDatabaseUpdateSag           = errors.New("error updating sager in database")
	ErrDatabaseUpdateSagstrin      = errors.New("error updating sagstrin in database")
	ErrDatabaseUpdateSagstrinstype = errors.New("error updating sagstrinstype in database")
)

type dbRepository struct {
	db *gorm.DB
}

func newDatabaseRepo(host string) *dbRepository {
	db, err := gorm.Open(sqlite.Open(host), &gorm.Config{ /*Logger: logger.Default.LogMode((logger.Silent))*/ })
	if err != nil {
		panic(errors.Join(ErrDatabaseConnection, err))
	}

	db.AutoMigrate(&ftoda.Sag{}, &ftoda.Emneord{}, &ftoda.EmneordSag{}, &ftoda.Sagstrin{})
	return &dbRepository{
		db: db,
	}
}

/*
	Sag
*/

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

func (repo *dbRepository) getSagerByTypeWithSagstrin(sagtype int, limit int) (sager []ftoda.SagSagstrin, err error) {

	/*

		query := db.Table("order").Select("MAX(order.finished_at) as latest").Joins("left join user user on order.user_id = user.id").Where("user.age > ?", 18).Group("order.user_id")
		db.Model(&Order{}).Joins("join (?) q on order.finished_at = q.latest", query).Scan(&results)
		// SELECT `order`.`user_id`,`order`.`finished_at` FROM `order` join (SELECT MAX(order.finished_at) as latest FROM `order` left join user user on order.user_id = user.id WHERE user.age > 18 GROUP BY `order`.`user_id`) q on order.finished_at = q.latest


	*/

	result := repo.db.Where("typeid = ?", sagtype).Find(&sager)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return sager, errors.Join(ErrRepoGettingSag, err)
	}

	return sager, err
}

// Done with test
func (repo *dbRepository) updateSager(sager []ftoda.Sag) (int64, error) {
	result := repo.db.Save(&sager)
	if result.Error != nil {
		return result.RowsAffected, errors.Join(ErrDatabaseUpdateSag, result.Error)
	}

	return result.RowsAffected, result.Error
}

/*
	Sagstrin
*/

func (repo *dbRepository) getSagstrinById(id int) (sagstrin ftoda.Sagstrin, err error) {
	result := repo.db.First(&sagstrin, id)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return sagstrin, errors.Join(ErrRepoGettingSagstrin, err)
	}

	return sagstrin, err
}

// Done with test
func (repo *dbRepository) getSagstrinBySagId(sagid int) (sagstrin []ftoda.Sagstrin, err error) {

	result := repo.db.Where("sag_id = ?", sagid).Find(&sagstrin)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return sagstrin, errors.Join(ErrRepoGettingSag, err)
	}

	return sagstrin, err
}

func (repo *dbRepository) updateSagstrin(sagstrin []ftoda.Sagstrin) (rows int64, err error) {
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&sagstrin)
	if result.Error != nil {
		return rows, errors.Join(ErrDatabaseUpdateSagstrin, result.Error)
	}
	return rows, err
}

/*
	Sagstrinstype
*/

func (repo *dbRepository) getSagstrinstype() (sagstrintypes []ftoda.Sagstrinstype, err error) {
	result := repo.db.Find(&sagstrintypes)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return sagstrintypes, errors.Join(ErrRepoGettingSagstrinstype, err)
	}
	return sagstrintypes, err
}

func (repo *dbRepository) updateSagstrintype(sagstrintypes []ftoda.Sagstrinstype) (rows int64, err error) {
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&sagstrintypes)
	if result.Error != nil {
		return rows, errors.Join(ErrDatabaseUpdateSagstrinstype, result.Error)
	}

	return rows, err
}

/*
	Emneord
*/

func (repo *dbRepository) getEmneordBySagId(sagid int) (emneord []ftoda.Emneord, err error) {
	/*
		var emneordsag []ftoda.EmneordSag
		//result := repo.db.Where("sagid = ?", sagid).Find(&emneordsag)

		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return emneord, errors.Join(ErrRepoGettingSag, err)
		}
		//db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
		return emneord, err
	*/

	return
}

func (repo *dbRepository) updateEmneord(emneord []ftoda.Emneord) (rows int64, err error) {
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&emneord)
	if result.Error != nil {
		return rows, errors.Join(ErrDatabaseUpdateSagstrin, result.Error)
	}
	return result.RowsAffected, err
}

func (repo *dbRepository) updateEmneordSag(emneordSager []ftoda.EmneordSag) (rows int64, err error) {
	result := repo.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&emneordSager)
	if result.Error != nil {
		return rows, errors.Join(ErrDatabaseUpdateSagstrin, result.Error)
	}
	return result.RowsAffected, err
}
