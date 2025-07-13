package lovforslag

import (
	"fmt"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
)

type Lovforslag struct {
	ftoda.Sag
	Emneord []string //we need to add emneord
}

type Lovtrin struct {
	Type     string
	Sagstrin []ftoda.Sagstrin
}

func NewFromSagId(sagid int, repo *repository.FTODAService) (l Lovforslag, err error) {

	sag, err := repo.GetSagById(sagid)
	if err != nil {
		return l, err
	}

	emneord, err := repo.GetEmneordBySagId(sagid)
	if err != nil {
		return l, err
	}

	var emner []string
	for _, emo := range emneord {
		emner = append(emner, emo.Emneord)
	}

	fmt.Println(emneord)
	l = Lovforslag{
		Sag:     sag,
		Emneord: emner,
	}
	//then we add emneord

	//get sag, then do something right?

	//but this exposes the query model. I think the other should be different.

	return l, err
}
