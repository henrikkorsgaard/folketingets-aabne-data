package repository

import (
	"errors"
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
	/*
		q := odataQuery{
			entity: "Sag",
			filter: "typeid eq " + strconv.Itoa(sagtype),
			skip:   offset,
			top:    limit,
		}

		err = s.api.getData(q, &sager)
		if err != nil {
			return nil, errors.Join(ErrGettingSag, err)
		}
	*/
	return
}

/*
	Afstemning
*/

func (s *FTODAService) GetAfstemningBySagstrinId(sagstrinid int) (afstemning ftoda.Afstemning, err error) {
	/*
		q := odataQuery{
			entity: "Afstemning",
			filter: "sagstrinid eq " + strconv.Itoa(sagstrinid),
		}

		var afstemninger []Afstemning
		err = s.api.getData(q, &afstemninger)
		if err != nil {
			return afstemning, errors.Join(ErrGettingAfstemning, err)
		}*/

	return
}

/*
	Sagstrin
*/

func (s *FTODAService) GetSagstrinBySagsId(sagid int) (sagstrin []ftoda.Sagstrin, err error) {
	/*
		q := odataQuery{
			entity: "Sagstrin",
			filter: "sagid eq " + strconv.Itoa(sagid),
			order:  "asc",
			expand: "Sagstrinstype,Afstemning",
		}

		err = s.api.getData(q, &sagstrin)
		if err != nil {
			return sagstrin, errors.Join(ErrGettingSagstrin, err)
		}
	*/
	return
}

func (s *FTODAService) GetSagstrinById(id int) (sag ftoda.Sagstrin, err error) {
	/*
		q := odataQuery{
			entity: "Sagstrin",
			filter: "id eq " + strconv.Itoa(id),
			expand: "Sagstrinstype,Afstemning,Dagsordenspunkt,SagstrinAktør,SagstrinDokument",
		}

		var sagstrin []Sagstrin
		err = s.api.getData(q, &sagstrin)
		if err != nil {
			return sag, errors.Join(ErrGettingSagstrin, err)
		}
	*/
	return
}

/*
	Emneord
*/

func (s *FTODAService) GetEmneordById(id int) (emne ftoda.Emneord, err error) {
	/*
		q := odataQuery{
			entity: "Emneord",
			filter: "id eq " + strconv.Itoa(id),
		}

		var emner []Emneord
		err = s.api.getData(q, &emner)
		if err != nil {
			return emne, errors.Join(ErrGettingEmneord, err)
		}
	*/
	return
}

/*
	Lovforslag
*/

func (s *FTODAService) UpdateSagerByType(sagtype int) (sager []ftoda.Sag, updateCount int64, err error) {
	/*
		q := odataQuery{
			entity: "Sag",
			filter: "typeid eq " + strconv.Itoa(sagtype),
			skip:   0,
		}

		err = s.api.getData(q, &sager)
		if err != nil {
			return sager, updateCount, errors.Join(ErrGettingSag, err)
		}

		updateCount = s.db.insertBulk(sager)
	*/
	return
}

func (s *FTODAService) GetLovforslagCount() (count int64) {
	/*
		affectedRows := s.db.getRowCount("sags")
	*/
	return
}
