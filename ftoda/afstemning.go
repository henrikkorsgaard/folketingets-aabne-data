package ftoda

import (
	"fmt"
	"slices"
	"context"
	"sync"

	dataloader "github.com/graph-gophers/dataloader/v7"
)

var (
	afstemningLoader     *dataloader.Loader[int, *Afstemning]
	afstemningLoaderOnce sync.Once
)

func afstemningBatchFunction(ctx context.Context, keys []int) (results []*dataloader.Result[*Afstemning]) {
	repo := newRepository()
	afstemninger, err := repo.getAfstemningerByIds(keys)
	// This is database errors, we handle discrepencies below
	if err != nil {
		panic(err)
	}

	
	//Two error cases:
	// 1: Len(keys) == 1; len(afstemninger) == 0
	// We never iterate through the afstemninger and never find out we have an error
	// 2: Len(keys) > len(afstemninger)
	// We have a couple of errors that need handling
	
	// Instead of iterating through the database result,
	// We iterate through the keys and then use the slice.IndexFunc to see if an Afstemning with the key is in the results. If not, we append an error to the results.
	for _, key := range keys {
		i := slices.IndexFunc(afstemninger, func(afstemning Afstemning) bool {
			return afstemning.Id == key
		})

		if i == -1 {
			e := fmt.Errorf("record not found: Afstemning with id %d not found", key)
			results = append(results, &dataloader.Result[*Afstemning]{Data: &Afstemning{},Error: e})
		} else {
			results = append(results, &dataloader.Result[*Afstemning]{Data: &afstemninger[i]})
		}
	}
	
	return
}

func newAfstemmeLoader() *dataloader.Loader[int, *Afstemning] {
	afstemningLoaderOnce.Do(func() {
		cache := &dataloader.NoCache[int, *Afstemning]{}
		afstemningLoader = dataloader.NewBatchedLoader(afstemningBatchFunction, dataloader.WithCache[int, *Afstemning](cache))
	})

	return afstemningLoader
}

type Afstemning struct {
	Id              int `gorm:"primaryKey"`
	Nummer          int
	Konklusion      string
	Vedtaget        int
	Kommentar       string
	MødeId          int
	Type            string
	SagstringId     int
	Opdateringsdato string
}

func (Afstemning) TableName() string {
	return "Afstemning"
}

func LoadAfstemning(id int) (afstemning Afstemning, err error) {
	
	loader := newAfstemmeLoader()
	thunk := loader.Load(context.Background(), id)
	result, err := thunk()
	afstemning = *result
	
	return
}

func LoadAfstemninger(limit int, offset int) (afstemninger []Afstemning, err error) {
	//This should just load from the database directly
	repo := newRepository()
	return repo.getAfstemninger(limit, offset)
}

func (r *Repository) getAfstemninger(limit int, offset int) (afstemninger []Afstemning, err error) {
	result := r.db.Limit(limit).Offset(offset).Find(&afstemninger)
	err = result.Error
	return
}

func (r *Repository) getAfstemningerByIds(ids []int) (afstemninger []Afstemning, err error) {
	result := r.db.Table("Afstemning").Select("Afstemning.*, Afstemningstype.type").Joins("left join Afstemningstype on Afstemning.typeid = Afstemningstype.id").Find(&afstemninger, ids)

	err = result.Error
	
	return
}
