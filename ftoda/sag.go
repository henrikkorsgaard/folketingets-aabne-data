package ftoda 

import (
	"fmt"
	"slices"
	"context"
	"sync"

	dataloader "github.com/graph-gophers/dataloader/v7"
)

var (
	sagLoader *dataloader.Loader[int, *Sag]
	sagLoaderOnce sync.Once
)

func sagBatchFunction(ctx context.Context, keys[]int) (results []*dataloader.Result[*Sag]) {

	repo := newRepository()
	sager, err := repo.getSagerByIds(keys)
	if err != nil {
		panic(err)
	}

	// See note on afsteming.go
	for _, key := range keys {
		i := slices.IndexFunc(sager, func(sag Sag) bool {
			return sag.Id == key
		})

		if i == -1 {
			e := fmt.Errorf("record not found: Sag with id %d not found", key)
			results = append(results, &dataloader.Result[*Sag]{Data: &Sag{},Error: e})
		} else {
			results = append(results, &dataloader.Result[*Sag]{Data: &sager[i]})
		}
	}

	return 
}

func newSagLoader() *dataloader.Loader[int, *Sag] {
	sagLoaderOnce.Do(func(){
		cache := &dataloader.NoCache[int, *Sag]{}
		sagLoader = dataloader.NewBatchedLoader(sagBatchFunction, dataloader.WithCache[int, *Sag](cache))
	})

	return sagLoader
}


type Sag struct {
	Id int `gorm:"primaryKey"`
	Titel string
	TitelKort string `gorm:"column:titelkort"`
	Offentlighedskode string `gorm:"column:offentlighedskode"`
	Nummer string 
	NummerPrefix string `gorm:"column:nummerprefix"`
	NummerNumerisk string `gorm:"column:nummernumerisk"`
	NummerPostfix string `gorm:"column:nummerpostfix"`
	Resume string
	Afstemningskonklusion string `gorm:"column:afstemningskonklusion"`
	PeriodeId int 
	AfgørelsesResultatKode string `gorm:"column:afgørelsesresultatkode"`
	Baggrundsmateriale string 
	Opdateringsdato string
	StatsbudgetSag int 
	Begrundelse string
	Paragrafnummer int
	Paragraf string
	AfgørelsesDato string 
	Afgørelse string 
	RådsmødeDato string 
	Lovnummer string 
	LovnummerDato string
	Retsinformationsurl string 
	FremsatUnderSagId int 
	DeltUnderSagId int 

	// Foreign types
	Type string //Table Sagtype 
	Kategori string //Table Sagkategori
	Status string  //Table Sagsstatus
}

func (Sag) TableName() string {
	return "Sag"
}

func LoadSager(limit, offset int) (sager []Sag, err error) {
	repo := newRepository()
	return repo.getSager(limit, offset)
}

func LoadSag(id int) (sag Sag, err error) {
	loader := newSagLoader()
	thunk := loader.Load(context.Background(), id)
	result, err := thunk()
	sag = *result
	return	
}

func LoadSagerByType(limit, offset int, sagType string) (sager []Sag, err error ){
	repo := newRepository()
	return repo.getSagerByType(limit, offset, sagType)
}

//if this is part of a list of afstemninger we need this to be accommodated in the loader
func LoadSagBySagstrin(id int) (sag Sag, err error) {
	repo := newRepository()
	return repo.getSagBySagstrinId(id)
}

func (r *Repository) getSagBySagstrinId(id int) (sag Sag, err error) {
	result := r.db.Table("Sag").Select("Sag.*").Joins("left join Sagstrin on Sagstrin.sagid = Sag.id").Where("Sagstrin.id = ?",id).Find(&sag)
	err = result.Error

	return
}

func (r *Repository) getSager(limit, offset int) (sager []Sag, err error){
	result := r.db.Table("Sag").Limit(limit).Offset(offset).Select("Sag.*, Sagstype.type, Sagskategori.kategori, Sagsstatus.status").Joins("left join Sagstype on Sag.typeid = Sagstype.id").Joins("left join Sagskategori on Sag.kategoriid = Sagskategori.id").Joins("left join Sagsstatus on Sag.statusid = Sagsstatus.id").Find(&sager)
	err = result.Error
	return
}

func (r *Repository) getSagerByIds(ids []int) (sager []Sag, err error){
	result := r.db.Table("Sag").Select("Sag.*, Sagstype.type, Sagskategori.kategori, Sagsstatus.status").Joins("left join Sagstype on Sag.typeid = Sagstype.id").Joins("left join Sagskategori on Sag.kategoriid = Sagskategori.id").Joins("left join Sagsstatus on Sag.statusid = Sagsstatus.id").Find(&sager, ids)
	err = result.Error

	return
}

func (r *Repository) getSagerByType(limit, offset int, sagType string) (sager []Sag, err error){
	result := r.db.Table("Sag").Select("Sag.*, Sagstype.type, Sagskategori.kategori, Sagsstatus.status").Joins("left join Sagstype on Sag.typeid = Sagstype.id").Joins("left join Sagskategori on Sag.kategoriid = Sagskategori.id").Joins("left join Sagsstatus on Sag.statusid = Sagsstatus.id").Where("Sagstype.type = ?", sagType).Find(&sager)
	err = result.Error
	return 
}
