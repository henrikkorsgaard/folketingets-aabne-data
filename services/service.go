package ftoda

import (
	"errors"
	"strconv"
)

var (
	ErrGettingSagstrin   = errors.New("error getting sagstrin")
	ErrGettingAfstemning = errors.New("error getting afstemning")
	ErrGettingEmneord    = errors.New("error getting emneord")
	ErrGettingSag        = errors.New("error getting sag")
)

// Todo - rename to store
type FTODAService struct {
	api *apiRepository
	db  *dbRepository
}

func NewFTODAService(odaHost string, dbHost string) FTODAService {

	// Host should come from either a factory or .env
	repo := newAPIRepository(odaHost)
	db := newDBRepository(dbHost)
	return FTODAService{
		api: repo,
		db:  db,
	}
}

/*
	Afstemning
*/

func (s *FTODAService) GetAfstemningBySagstrinId(sagstrinid int) (afstemning Afstemning, err error) {

	q := odataQuery{
		entity: "Afstemning",
		filter: "sagstrinid eq " + strconv.Itoa(sagstrinid),
	}

	var afstemninger []Afstemning
	err = s.api.getData(q, &afstemninger)
	if err != nil {
		return afstemning, errors.Join(ErrGettingAfstemning, err)
	}

	return afstemninger[0], nil
}

/*
	Sagstrin
*/

func (s *FTODAService) GetSagstrinBySagsId(sagid int) (sagstrin []Sagstrin, err error) {

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

	return sagstrin, nil
}

func (s *FTODAService) GetSagstrinById(id int) (sag Sagstrin, err error) {

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

	return sagstrin[0], nil
}

/*
	Emneord
*/

func (s *FTODAService) GetEmneordById(id int) (emne Emneord, err error) {

	q := odataQuery{
		entity: "Emneord",
		filter: "id eq " + strconv.Itoa(id),
	}

	var emner []Emneord
	err = s.api.getData(q, &emner)
	if err != nil {
		return emne, errors.Join(ErrGettingEmneord, err)
	}

	return emner[0], nil
}

/*
	Lovforslag
*/

func (s *FTODAService) GetSagById(id int) (sag Sag, err error) {

	q := odataQuery{
		entity: "Sag",
		filter: "typeid eq 3 and id eq " + strconv.Itoa(id),
		expand: "EmneordSag",
	}

	var sager []Sag
	err = s.api.getData(q, &sager)
	if err != nil {
		return sag, errors.Join(ErrGettingSag, err)
	}
	return sager[0], nil
}

func (s *FTODAService) GetSagerByType(sagtype int, limit int, offset int) (sager []Sag, err error) {

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

	return sager, nil
}

func (s *FTODAService) UpdateSagerByType(sagtype int) (sager []Sag, updateCount int64, err error) {

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

	return sager, updateCount, nil
}

func (s *FTODAService) GetLovforslagCount() int64 {

	affectedRows := s.db.getRowCount("sags")
	return affectedRows
}
