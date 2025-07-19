package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestServiceFetchAndPersistToDatabase(t *testing.T) {
	defer cleanup()

	is := is.New(t)
	sagid := 102467
	s := New("oda.ft.dk", testdb)
	sag, err := s.GetSagById(sagid)
	is.NoErr(err)
	is.Equal(sag.Id, sagid) // not neccesary, but still

	dbsag, err := s.db.getSagById(sagid)
	is.NoErr(err)
	is.Equal(dbsag.Id, sagid)

	sagstrin, err := s.db.getSagstrinBySagId(sagid)
	is.NoErr(err)
	is.True(len(sagstrin) > 0)
	if len(sagstrin) > 0 {
		is.Equal(sag.Sagstrin[0].Id, sagstrin[0].Id)
	}
}

func cleanup() {
	fmt.Println("Removing test database")
	err := os.Remove(testdb)
	if err != nil {
		panic(err)
	}
}
