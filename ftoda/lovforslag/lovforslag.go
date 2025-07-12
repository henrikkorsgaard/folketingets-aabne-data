package lovforslag

import (
	ftoda "github.com/henrikkorsgaard/folketingets-aabne-data/services"
)

type Lovforslag struct {
	ftoda.Sag //We want to refactor this I assume
}

func New(sagid int) (l Lovforslag, err error) {

	//get sag, then do something right?

	//but this exposes the query model. I think the other should be different.

	return l, err
}
