package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/matryer/is"
)

var testdb = "test.db"

// TODO: I want to remove the database after the tests
func TestSagAndSagstrinRelation(t *testing.T) {
	defer cleanup()

	is := is.New(t)

	st1 := ftoda.Sagstrin{
		Id:    1,
		SagId: 1,
	}

	st2 := ftoda.Sagstrin{
		Id:    2,
		SagId: 1,
	}

	st3 := ftoda.Sagstrin{
		Id:    3,
		SagId: 1,
	}

	sag := ftoda.Sag{
		Id:        1,
		Titel:     "Test Sag",
		TitelKort: "Test Sag",
		Sagstrin:  []ftoda.Sagstrin{st1, st2, st3},
	}

	db := newDatabaseRepo(testdb)
	r, err := db.updateSager([]ftoda.Sag{sag})
	is.NoErr(err)
	is.Equal(r, int64(1))

	sagstrin, err := db.getSagstrinBySagId(1)
	is.NoErr(err)
	is.Equal(sagstrin[0], st1)

}

func cleanup() {
	fmt.Println("Removing test database")
	err := os.Remove(testdb)
	if err != nil {
		panic(err)
	}
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
