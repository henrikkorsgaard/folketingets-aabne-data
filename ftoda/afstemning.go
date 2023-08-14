package ftoda

import (
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

	if err != nil {
		panic(err)
	}

	for _, afstemning := range afstemninger {
		results = append(results, &dataloader.Result[*Afstemning]{Data: &afstemning})
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
	result := r.db.Find(&afstemninger, ids)
	err = result.Error
	return
}
