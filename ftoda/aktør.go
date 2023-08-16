package ftoda

import (
	"fmt"
	"slices"
	"context"
	"sync"
	dataloader "github.com/graph-gophers/dataloader/v7"
)

var (
	aktørLoader *dataloader.Loader[int, *Aktør]
	aktørLoaderOnce sync.Once
) 

func aktørBatchFunction(ctx context.Context, keys[] int) (results []*dataloader.Result[*Aktør]) {
	repo := newRepository()
	aktører, err := repo.getAktørerByIds(keys)

	if err != nil {
		panic(err)
	}


	// Instead of iterating through the database result,
	// We iterate through the keys and then use the slice.IndexFunc to see if an Afstemning with the key is in the results. If not, we append an error to the results.
	for _, key := range keys {
		i := slices.IndexFunc(aktører, func(aktør Aktør) bool {
			return aktør.Id == key
		})

		if i == -1 {
			e := fmt.Errorf("record not found: Aktør with id %d not found", key)
			results = append(results, &dataloader.Result[*Aktør]{Data: &Aktør{},Error: e})
		} else {
			results = append(results, &dataloader.Result[*Aktør]{Data: &aktører[i]})
		}
	}

	return 
}

func newAktørLoader() *dataloader.Loader[int, *Aktør] {
	aktørLoaderOnce.Do(func() {
		cache := &dataloader.NoCache[int, *Aktør]{}
		aktørLoader = dataloader.NewBatchedLoader(aktørBatchFunction, dataloader.WithCache[int, *Aktør](cache))
	})

	return aktørLoader
}

type Aktør struct {
	Id int `gorm:"primaryKey"`
	Type string
	GruppeNavnKort string
	Navn string
	Fornavn string 
	Efternavn string
	Biografi string
	Periode int 
	Opdateringsdato string
	Startdato string
	Slutdato string
}

func (Aktør) TableName() string {
	return "Aktør"
}

func LoadAktørById(id int) (aktør Aktør, err error) {
	loader := newAktørLoader()
	thunk := loader.Load(context.Background(), id)

	result, err := thunk()

	aktør = *result

	return
}

func LoadAktørByName(name string) (aktør Aktør, err error) {
	repo := newRepository()
	return repo.getAktørByName(name)

	return
}

func LoadAktører(limit int, offset int) (aktører []Aktør, err error) {
	repo := newRepository()
	return repo.getAktører(limit, offset)
}

func LoadAktørerByType(limit int, offset int, aktørType string) (aktører []Aktør, err error) {
	repo := newRepository()
	return repo.getAktørerByType(limit, offset, aktørType)
}

func (r *Repository) getAktører(limit int, offset int) (aktører []Aktør, err error) {
	result := r.db.Table("Aktør").Limit(limit).Offset(offset).Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Find(&aktører)
	err = result.Error
	return
}

func (r *Repository) getAktørerByType(limit int, offset int, aktørType string) (aktører []Aktør, err error) {
	result := r.db.Table("Aktør").Limit(limit).Offset(offset).Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Where("Aktørtype.type = ?", aktørType).Find(&aktører)
	err = result.Error
	return
}

func (r *Repository) getAktørerByIds(ids []int) (aktører []Aktør, err error) {
	result := r.db.Table("Aktør").Select("Aktør.*, Aktørtype.type").Joins("left join Aktørtype on Aktør.typeid = Aktørtype.id").Find(&aktører, ids)
	err = result.Error
	return 
}

func (r *Repository) getAktørByName(name string) (aktør Aktør, err error) {
	result := r.db.Where("navn = ?",name).Find(&aktør)
	err = result.Error
	return 
}
