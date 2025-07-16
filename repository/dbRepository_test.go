package repository

import (
	"testing"

	"github.com/matryer/is"
)

// TODO: I want to remove the database after the tests
func TestGetSagerByTypeWithSagstrin(t *testing.T) {
	is := is.New(t)
	s := New("oda.ft.dk", "test.db")
	//prep insert
	s.GetSagerByTypeWithSagstrin(3, 5, 0)
	sager, err := s.db.getSagerByTypeWithSagstrin(3, 5)
	is.NoErr(err)
	is.Equal(len(sager), 5)
}
