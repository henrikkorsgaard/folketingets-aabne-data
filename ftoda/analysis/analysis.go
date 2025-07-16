package analysis

import (
	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
)

// what should this return?
// the template or the data

// we can get:
// Distribution across types
//

func LovforslagSagstrinTypeDistribution(lovforslagCount int, repo *repository.FTODAService) {

	// we also need sagstrin status to be able to match this.

	_, err := repo.GetSagerByTypeWithSagstrin(3, 200, 0)
	if err != nil {
		panic(err) //just for now
	}

	//lets get 100 lovforslag
	//expand with sagstrin for fast query or individual queries?
	//We only need it for this analysis.
	//but without some bulk work, e.g. getsagstrinbyids []sagid, then  it will be 1+100 queries.
	// with the raw query, I can use it as an update point.
	// that will be an issue if I do not implement releational types. What the heck.
}
