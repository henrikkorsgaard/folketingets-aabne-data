package repository

import (
	"testing"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/matryer/is"
)

var testdb = "test.db"

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
