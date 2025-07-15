package repository

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
)

var (
	ErrGettingSagstrin   = errors.New("error getting sagstrin")
	ErrGettingAfstemning = errors.New("error getting afstemning")
	ErrGettingEmneord    = errors.New("error getting emneord")
	ErrGettingSag        = errors.New("error getting sag")
)

// Todo - rename to store
type FTODAService struct {
	api *odataRepository
	db  *dbRepository
}

func New(odaHost string, dbHost string) FTODAService {

	return FTODAService{
		api: newOdataRepo(odaHost),
		db:  newDatabaseRepo(dbHost),
	}
}

/*
	Sag
*/

func (s *FTODAService) GetSagById(id int) (sag ftoda.Sag, err error) {

	sag, err = s.db.getSagById(id)
	if err != nil {
		return sag, err
	}

	//guard function returns if db returns sag
	if !reflect.ValueOf(sag).IsZero() {
		return sag, err
	}

	sag, err = s.api.getSagById(id)
	if err != nil {
		return sag, err
	}

	_, err = s.db.updateSager([]ftoda.Sag{sag})
	if err != nil {
		return sag, err
	}

	return sag, err
}

func (s *FTODAService) GetSagerByType(sagtype int, limit int, offset int) (sager []ftoda.Sag, err error) {

	sager, err = s.db.getSagerByType(sagtype)
	if err != nil {
		return sager, err
	}

	if len(sager) > 0 {
		return sager, err
	}

	sager, err = s.api.getSagerByType(sagtype)
	if err != nil {
		return sager, err
	}

	_, err = s.db.updateSager(sager)
	if err != nil {
		return sager, err
	}

	return sager, err
}

func (s *FTODAService) GetSagerByTypeWithSagstrin(sagtype int, limit int, offset int) (sager []ftoda.SagSagstrin, err error) {
	/*
		sager, err = s.db.getSagerByTypeWithSagstrin(sagtype)
		if err != nil {
			return sager, err
		}

		if len(sager) > 0 {
			return sager, err
		}*/

	sager, err = s.api.getSagerByTypeWithSagstrin(sagtype)
	if err != nil {
		return sager, err
	}
	/*
		_, err = s.db.updateSagerWithSagstrin(sager)
		if err != nil {
			return sager, err
		}*/

	return sager, err
}

/*
	Sagstrin
*/

func (s *FTODAService) GetSagstrinBySagsId(sagid int) (sagstrin []ftoda.Sagstrin, err error) {

	sagstrin, err = s.db.getSagstrinBySagId(sagid)
	if err != nil {
		return sagstrin, err
	}

	if len(sagstrin) > 0 {
		return sagstrin, err
	}

	sagstrin, err = s.api.getSagstrinBySagId(sagid)
	if err != nil {
		return sagstrin, err
	}

	_, err = s.db.updateSagstrin(sagstrin)
	if err != nil {
		return sagstrin, err
	}

	return sagstrin, err
}

func (s *FTODAService) GetSagstrinById(id int) (sagstrin ftoda.Sagstrin, err error) {

	sagstrin, err = s.db.getSagstrinById(id)
	if err != nil {
		return sagstrin, err
	}

	//guard function returns if db returns sag
	if !reflect.ValueOf(sagstrin).IsZero() {
		return sagstrin, err
	}

	sagstrin, err = s.api.getSagstrinById(id)
	if err != nil {
		return sagstrin, err
	}

	_, err = s.db.updateSagstrin([]ftoda.Sagstrin{sagstrin})
	if err != nil {
		return sagstrin, err
	}

	return sagstrin, err
}

/*
	Emneord
*/

func (s *FTODAService) GetEmneordBySagId(sagid int) (emneord []ftoda.Emneord, err error) {

	emneord, err = s.api.getEmneordBySagId(sagid)
	if err != nil {
		return emneord, err
	}

	var emneordsager []ftoda.EmneordSag
	for _, emn := range emneord {
		emneordsager = append(emneordsager, emn.EmneordSag)
	}

	r1, err := s.db.updateEmneordSag(emneordsager)
	if err != nil {
		return emneord, err
	}

	fmt.Printf("Emneordsag upsert rows %d\n", r1)

	r2, err := s.db.updateEmneord(emneord)
	if err != nil {
		return emneord, err
	}

	fmt.Printf("Emneord upsert rows %d\n", r2)

	return emneord, err
}
