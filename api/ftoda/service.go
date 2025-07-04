package ftoda

import (
	"encoding/json"
	"fmt"
	"strconv"
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

	fmt.Println(q.PrettyUrl(s.api.host))
	odata, err := s.api.getData(q)
	if err != nil {
		fmt.Printf("error from afstemning: %s\n", err)
		return afstemning, err
	}

	var afstemninger []Afstemning
	err = json.Unmarshal(odata.Result, &afstemninger)
	if err != nil {
		fmt.Printf("error from afstemning: %s\n", err)
		return afstemning, err
	}

	return afstemninger[0], nil
}

/*
	Sagstrin
*/

func (s *FTODAService) GetSagstrinBySagsId(sagid int) (sagstrin []Sagstrin, err error) {
	//First we should check a database, but that is not created yet
	//If not found in database, then we get it from the api

	q := odataQuery{
		entity: "Sagstrin",
		filter: "sagid eq " + strconv.Itoa(sagid),
	}

	fmt.Println(q.PrettyUrl(s.api.host))
	// this need to be moved into a different repo service
	odata, err := s.api.getData(q)
	if err != nil {
		fmt.Printf("error from GetSagstrinBySagsId: %s\n", err)
		return sagstrin, err
	}

	// this need to be moved into a different repo service
	err = json.Unmarshal(odata.Result, &sagstrin)
	if err != nil {
		fmt.Printf("error from  GetSagstrinBySagsId: %s\n", err)
		return sagstrin, err
	}

	return sagstrin, nil
}

/*
	Lovforslag
*/

func (s *FTODAService) GetLovforslagById(id int) (sag Sag, err error) {
	//First we should check a database, but that is not created yet
	//If not found in database, then we get it from the api

	q := odataQuery{
		entity: "Sag",
		filter: "typeid eq 3 and id eq " + strconv.Itoa(id),
	}
	// this need to be moved into a different repo service
	odata, err := s.api.getData(q)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return sag, err
	}

	// this need to be moved into a different repo service
	var sager []Sag
	err = json.Unmarshal(odata.Result, &sager)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return sag, err
	}

	return sager[0], nil
}

// offset map into skip next for now
func (s *FTODAService) GetLovforslag(limit int, offset int) ([]Sag, error) {

	q := odataQuery{
		entity: "Sag",
		filter: "typeid eq 3",
		skip:   offset,
		top:    limit,
	}

	odata, err := s.api.getData(q)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return nil, err
	}

	var sager []Sag
	err = json.Unmarshal(odata.Result, &sager)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return nil, err
	}

	return sager, nil
}

func (s *FTODAService) UpdateLovforslag() ([]Sag, int64, error) {

	q := odataQuery{
		entity: "Sag",
		filter: "typeid eq 3",
		skip:   0,
	}

	odata, err := s.api.getData(q)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return nil, 0, err
	}

	var sager []Sag
	err = json.Unmarshal(odata.Result, &sager)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return nil, 0, err
	}

	affectedRows := s.db.insertBulk(sager)

	return sager, affectedRows, nil
}

func (s *FTODAService) GetLovforslagCount() int64 {

	affectedRows := s.db.getRowCount("sags")
	return affectedRows
}
