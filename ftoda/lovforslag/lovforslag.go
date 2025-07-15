package lovforslag

import (
	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
)

type Lovforslag struct {
	ftoda.Sag
	Emneord []string //we need to add emneord
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

	l = Lovforslag{
		Sag:     sag,
		Emneord: emner,
	}

	return l, err
}
