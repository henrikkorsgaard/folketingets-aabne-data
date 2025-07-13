package repository

import (
	"errors"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
)

var (
	ErrRepoGettingSagstrin   = errors.New("error getting sagstrin from repo")
	ErrRepoGettingAfstemning = errors.New("error getting afstemning from repo")
	ErrRepoGettingEmneord    = errors.New("error getting emneord from repo")
	ErrRepoGettingSag        = errors.New("error getting sag from repo")
)

type Repository interface {
	getSagById(id int) (sag ftoda.Sag, err error)
	getSagerByType(sagtype int) (sager []ftoda.Sag, err error)
}
