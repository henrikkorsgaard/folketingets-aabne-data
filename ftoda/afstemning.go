package ftoda

import (

	//"gorm.io/gorm"
	//graphql "github.com/graph-gophers/graphql-go"
)

type Afstemning struct {
	Id int32 `gorm:"primaryKey"`
	Nummer int32
	Konklusion string
	Vedtaget int32
	Kommentar string
	MødeID int32 
	Type string
	SagstrinID int32 
	//Opdateringsdato graphql.Time // need to implement scanner before I can use this.
}

func (Afstemning) TableName() string {
	return "Afstemning"
}

// Return one 
func (r *Repository) GetAfstemning(id int32) (afstemning Afstemning, err error) {
	r.db.First(&afstemning)
	return 
}

// Return all * with the provided limits and offsets
func (r *Repository) GetAllAfstemning() (afstemninger []Afstemning, err error) {
	return
}

// Return ids --> only used internally with the loader
func (r *Repository) GetAfstemningList(ids []int32) (afsemninger []Afstemning, err error) {
	return
}