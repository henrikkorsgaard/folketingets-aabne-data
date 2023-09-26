package ftoda

import (
	"context"
	"sync"

	dataloader "github.com/graph-gophers/dataloader/v7"
)

var (
	aktorRelationLoader *dataloader.Loader[int, *[]AktorRelation]
	aktorRelationLoaderOnce sync.Once
)

func aktorRelationBatchFunction(ctx context.Context, keys []int)(results []*dataloader.Result[*[]AktorRelation]) {
	repo := newRepository()
	
	aktorRelationer, err :=  repo.getRelations(keys)
	
	if err != nil {
		panic(err)
	}

	aktorRelationsByKey := make(map[int][]AktorRelation)

	for _, ar := range aktorRelationer {
		aktorRelationsByKey[ar.FraAktorId] = append(aktorRelationsByKey[ar.FraAktorId], ar)
	}

	for _, key := range keys {
		relationer := aktorRelationsByKey[key]
		results = append(results, &dataloader.Result[*[]AktorRelation]{Data:&relationer})
	}

	return

}

func newAktorRelationLoader() *dataloader.Loader[int, *[]AktorRelation] {
	aktorRelationLoaderOnce.Do(func(){
		cache := &dataloader.NoCache[int, *[]AktorRelation]{}
		aktorRelationLoader = dataloader.NewBatchedLoader(aktorRelationBatchFunction, dataloader.WithCache[int, *[]AktorRelation](cache))
	})

	return aktorRelationLoader
}

type AktorRelation struct {
	Id int `gorm:"primaryKey"`
	FraAktorId int `gorm:"column:fraaktørid"`
	TilAktorId int `gorm:"column:tilaktørid"`
	Rolle string 
	TilAktorType string
	TilAktorNavn string
}

func (AktorRelation) TableName() string {
	return "AktørAktør"
}

func LoadAktorRelations(id int) (aktorRelations []AktorRelation, err error) {

	loader := newAktorRelationLoader()
	thunk := loader.Load(context.Background(), id)

	result, err := thunk()

	aktorRelations = *result 

	return 
}

func (r *Repository) getRelations(ids []int) (aktorRelations []AktorRelation, err error) {
	result := r.db.Table("AktørAktør").Select("AktørAktør.id, AktørAktør.fraaktørid, AktørAktør.tilaktørid, AktørAktørRolle.rolle AS rolle, Aktør.navn AS TilAktorNavn, Aktør.typeid, AktørType.type AS TilAktorType").Joins("left join AktørAktørRolle on AktørAktørRolle.id = AktørAktør.rolleid").Joins("left join Aktør on AktørAktør.tilaktørid = Aktør.id").Joins("left join AktørType on AktørType.id = Aktør.typeid").Where("fraaktørid IN ?", ids).Find(&aktorRelations)
	err = result.Error
	return 
}
