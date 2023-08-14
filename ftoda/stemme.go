package ftoda

import (
	"context"
	"sync"

	dataloader "github.com/graph-gophers/dataloader/v7"
)

var (
	stemmeLoader *dataloader.Loader[int, *[]Stemme] // Could just be the batch loader
	stemmeLoaderOnce 	sync.Once
)

func StemmeBatchFunction(ctx context.Context, keys []int) (results []*dataloader.Result[*[]Stemme]) {
	repo := NewRepository()

	stemmer, err := repo.GetStemmeByAfstemningIds(keys)
	if err != nil {
		panic(err) // Want to force a solution if an error occurs.
	}
	
	stemmerByKey := make(map[int][]Stemme)

	for _, stemme := range stemmer {
		stemmerByKey[stemme.AfstemningId] = append(stemmerByKey[stemme.AfstemningId], stemme)
	}
	
	for _, key := range keys {
		stmr := stemmerByKey[key]
		results = append(results,  &dataloader.Result[*[]Stemme]{Data: &stmr})
	}

	return
}

func NewStemmeLoader() *dataloader.Loader[int, *[]Stemme] {
	stemmeLoaderOnce.Do(func(){
		cache := &dataloader.NoCache[int, *[]Stemme]{}
		stemmeLoader = dataloader.NewBatchedLoader(StemmeBatchFunction, dataloader.WithCache[int, *[]Stemme](cache))
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

func (r *Repository) GetStemme(id int) (stemme Stemme, err error) {
	result := r.db.First(&stemme, id)
	err = result.Error
	return
}

func (r *Repository) GetAllStemme(limit int, offset int) (stemmer []Stemme, err error) {
	result := r.db.Limit(limit).Offset(offset).Find(&stemmer)
	err = result.Error

	return
}

func (r *Repository) GetStemmeByIds(ids []int) (stemmer []Stemme, err error) {
	result := r.db.Find(&stemmer, ids)
	err = result.Error
	return
}

func (r *Repository) GetStemmeByAfstemningIds(ids []int) (stemmer []Stemme, err error) {
	result := r.db.Where("afstemningid IN ?", ids).Find(&stemmer)
	err = result.Error
	return
}
