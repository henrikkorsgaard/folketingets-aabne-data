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
	var stemmer []Stemme
	var err error
	if ctx.Value("parent") == "afstemning" {
		stemmer, err = repo.getStemmerByAfstemningIds(keys)
	}
	
	if ctx.Value("parent") == "aktor" {
		stemmer, err = repo.getStemmerByAktorIds(keys)
	}


	if err != nil {
		panic(err) // Want to force a solution if an error occurs.
	}

	stemmerByKey := make(map[int][]Stemme)

	for _, stemme := range stemmer {
		if ctx.Value("parent") == "afstemning" {
			stemmerByKey[stemme.AfstemningId] = append(stemmerByKey[stemme.AfstemningId], stemme)
		} 

		if ctx.Value("parent") == "aktor" {
			stemmerByKey[stemme.AktorId] = append(stemmerByKey[stemme.AktorId], stemme)
		}
	}

	// this does not handle errors atm.
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
	AktorId         int `gorm:"column:aktørid"`
	Opdateringsdato string
}

func (Stemme) TableName() string {
	return "Stemme"
}

func LoadStemmerFromAktor(id int) (stemmer []Stemme, err error) {
	loader := newStemmeLoader()

	ctx := context.WithValue(context.Background(), "parent", "aktor")
	thunk := loader.Load(ctx, id)

	result, err := thunk()

	stemmer = *result
	return
}

func LoadStemmerFromAfstemning(id int) (stemmer []Stemme, err error) {
	loader := newStemmeLoader()

	ctx := context.WithValue(context.Background(), "parent", "afstemning")
	thunk := loader.Load(ctx, id)

	result, err := thunk()

	stemmer = *result
	return
}

func (r *Repository) getStemmerByAfstemningIds(ids []int) (stemmer []Stemme, err error) {
	result := r.db.Select("Stemme.*, Stemmetype.type").Joins("left join Stemmetype on Stemme.typeid = Stemmetype.id").Where("afstemningid IN ?", ids).Find(&stemmer)
	err = result.Error
	return
}

func (r *Repository) getStemmerByAktorIds(ids []int) (stemmer []Stemme, err error) {
	result := r.db.Select("Stemme.*, Stemmetype.type").Joins("left join Stemmetype on Stemme.typeid = Stemmetype.id").Where("aktørid IN ?", ids).Find(&stemmer)
	err = result.Error
	return
}
