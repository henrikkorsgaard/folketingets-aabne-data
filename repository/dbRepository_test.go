package repository

import (
	"fmt"
	"testing"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/matryer/is"
)

// TODO: I want to remove the database after the tests
func TestSagAndSagstrinRelation(t *testing.T) {
	is := is.New(t)

	//insert individualle
	sag := ftoda.Sag{
		Id:        1,
		Titel:     "Test Sag",
		TitelKort: "Test Sag",
	}

	db := newDatabaseRepo("test.db")
	r, err := db.updateSager([]ftoda.Sag{sag})
	is.NoErr(err)
	fmt.Println(r)
}

/*
func TestGetSagerByTypeWithSagstrin(t *testing.T) {
	is := is.New(t)
	s := New("oda.ft.dk", "test.db")
	//prep insert
	s.GetSagerByTypeWithSagstrin(3, 5, 0)
	sager, err := s.db.getSagerByTypeWithSagstrin(3, 5)
	is.NoErr(err)
	is.Equal(len(sager), 5)
}*/
