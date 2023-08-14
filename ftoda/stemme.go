package ftoda

import (
	"context"
	"sync"

	dataloader "github.com/graph-gophers/dataloader/v7"
)

// this should be internal
var (
	stemmeLoader     *dataloader.Loader[int, *[]Stemme] // Could just be the batch loader
	stemmeLoaderOnce sync.Once
)

func stemmeBatchFunction(ctx context.Context, keys []int) (results []*dataloader.Result[*[]Stemme]) {
	repo := newRepository()
	stemmer, err := repo.getStemmeByAfstemningIds(keys)
	if err != nil {
		panic(err) // Want to force a solution if an error occurs.
	}

	stemmerByKey := make(map[int][]Stemme)

	for _, stemme := range stemmer {
		stemmerByKey[stemme.AfstemningId] = append(stemmerByKey[stemme.AfstemningId], stemme)
	}

	for _, key := range keys {
		stmr := stemmerByKey[key]
		results = append(results, &dataloader.Result[*[]Stemme]{Data: &stmr})
	}

	return
}

func newStemmeLoader() *dataloader.Loader[int, *[]Stemme] {
	stemmeLoaderOnce.Do(func() {
		cache := &dataloader.NoCache[int, *[]Stemme]{}
		stemmeLoader = dataloader.NewBatchedLoader(stemmeBatchFunction, dataloader.WithCache[int, *[]Stemme](cache))
	})

	return stemmeLoader
}

type Stemme struct {
	Id              int `gorm:"primaryKey"`
	Type            string
	AfstemningId    int `gorm:"column:afstemningid"`
	AktørId         int
	Opdateringsdato string
}

func (Stemme) TableName() string {
	return "Stemme"
}

func LoadStemmerFromAfstemning(id int) (stemmer []Stemme, err error) {
	loader := newStemmeLoader()
	thunk := loader.Load(context.Background(), id)

	result, err := thunk()

	stemmer = *result
	return
}

func (r *Repository) getStemmeByAfstemningIds(ids []int) (stemmer []Stemme, err error) {
	result := r.db.Where("afstemningid IN ?", ids).Find(&stemmer)
	err = result.Error
	return
}
