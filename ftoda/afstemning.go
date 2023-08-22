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
	MødeId          int `gorm:"column:mødeid"`
	Type            string
	SagstrinId     int
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

func LoadAfstemninger(limit, offset int) (afstemninger []Afstemning, err error) {
	//This should just load from the database directly
	repo := newRepository()
	return repo.getAfstemninger(limit, offset)
}

func LoadAfstemningerWithKommentar(limit, offset int) (afstemninger []Afstemning, err error) {
	repo := newRepository()
	return repo.getAfstemningerWhereFieldNotNull(limit, offset, "kommentar")
}

func LoadAfstemningerByType(limit, offset int, afstemningsType string) (afstemninger []Afstemning, err error) {
	//This should just load from the database directly
	repo := newRepository()
	return repo.getAfstemningerByType(limit, offset, afstemningsType)
}

func (r *Repository) getAfstemninger(limit, offset int) (afstemninger []Afstemning, err error) {
	result := r.db.Table("Afstemning").Limit(limit).Offset(offset).Select("Afstemning.*, Afstemningstype.type").Joins("left join Afstemningstype on Afstemning.typeid = Afstemningstype.id").Find(&afstemninger)
	err = result.Error
	return
}

func (r *Repository) getAfstemningerByType(limit, offset int, afstemningsType string) (afstemninger []Afstemning, err error) {
	result := r.db.Table("Afstemning").Limit(limit).Offset(offset).Select("Afstemning.*, Afstemningstype.type").Joins("left join Afstemningstype on Afstemning.typeid = Afstemningstype.id").Where("Afstemningstype.type = ?", afstemningsType).Find(&afstemninger)
	err = result.Error
	return
}

func (r *Repository) getAfstemningerWhereFieldNotNull(limit, offset int, field string) (afstemninger []Afstemning, err error) {
	result := r.db.Table("Afstemning").Limit(limit).Offset(offset).Select("Afstemning.*, Afstemningstype.type").Joins("left join Afstemningstype on Afstemning.typeid = Afstemningstype.id").Where("Afstemning." + field + " IS NOT NULL AND Afstemning." + field + " != ''").Find(&afstemninger)
	err = result.Error
	return
}

func (r *Repository) getAfstemningerByIds(ids []int) (afstemninger []Afstemning, err error) {
	result := r.db.Table("Afstemning").Select("Afstemning.*, Afstemningstype.type").Joins("left join Afstemningstype on Afstemning.typeid = Afstemningstype.id").Find(&afstemninger, ids)

	err = result.Error
	
	return
}
