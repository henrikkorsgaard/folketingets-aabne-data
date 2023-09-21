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

	var sager []Sag
	var err error 
	if ctx.Value("idtype") == "sagid" {
		sager, err = repo.getSagerByIds(keys)
	}

	if ctx.Value("idtype") == "sagstrinid" {
		sager, err = repo.getSagerBySagstrinIds(keys)
	}

	if err != nil {
		panic(err)
	}

	// See note on afsteming.go
	for _, key := range keys {
		i := slices.IndexFunc(sager, func(sag Sag) bool {
			// this is where we need to handle the key match
			if ctx.Value("idtype") == "sagid" {
				return sag.Id == key
			}

			if ctx.Value("idtype") == "sagstrinid" {
				return sag.SagstrinId == key
			}

			return false
		})

		if i == -1 {
			e := fmt.Errorf("record not found: Sag with %s %d not found",ctx.Value("idtype"), key)
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
	AfgorelsesResultatKode string `gorm:"column:afgorelsesresultatkode"`
	Baggrundsmateriale string 
	Opdateringsdato string
	StatsbudgetSag int 
	Begrundelse string
	Paragrafnummer int
	Paragraf string
	AfgorelsesDato string 
	Afgorelse string 
	RådsmodeDato string 
	Lovnummer string 
	LovnummerDato string
	Retsinformationsurl string 
	FremsatUnderSagId int 
	DeltUnderSagId int 

	// Foreign types
	Type string //Table Sagtype 
	Kategori string //Table Sagkategori
	Status string  //Table Sagsstatus
	SagstrinId int //Table Sagstrin, use-case when identifying sag by sagstrin (relation Afstemning)
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
	ctx := context.WithValue(context.Background(), "idtype", "sagid")
	thunk := loader.Load(ctx, id)
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
	loader := newSagLoader()
	ctx := context.WithValue(context.Background(), "idtype", "sagstrinid")
	thunk := loader.Load(ctx, id)
	result, err := thunk()
	sag = *result
	return	
}

func (r *Repository) getSagerBySagstrinIds(ids []int) (sager []Sag, err error) {
	result := r.db.Table("Sag").Select("Sag.*, Sagstrin.id AS SagstrinId").Joins("left join Sagstrin on Sagstrin.sagid = Sag.id").Where("Sagstrin.id IN ?",ids).Find(&sager)
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
