package repository

import (
	"testing"

	"github.com/matryer/is"
)

func TestSagstrintype(t *testing.T) {
	is := is.New(t)
	api := newOdataRepo("oda.ft.dk")
	sagstrinstype, err := api.getSagstrinstype()
	is.NoErr(err)
	is.Equal(len(sagstrinstype), 107)
}

func TestSagerByTypeWithSagstrin(t *testing.T) {
	is := is.New(t)
	api := newOdataRepo("oda.ft.dk")
	sager, err := api.getSagerByTypeWithSagstrin(3, 200)
	is.NoErr(err)
	is.Equal(len(sager), 200)
}
