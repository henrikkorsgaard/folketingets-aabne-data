package ftoda

import "fmt"

type AktorRelation struct {
	Id int `gorm:"primaryKey"`
	FraAktorId int `gorm:"column:fraaktørid"`
	TilAktorId int `gorm:"column:tilaktørid"`
	Rolle string 
	TilAktorNavn string
}

func (AktorRelation) TableName() string {
	return "AktørAktør"
}

func LoadAktorRelations(id int) (aktorRelations []AktorRelation, err error) {
	repo = newRepository()
	return repo.getRelations(id)
}

func (r *Repository) getRelations(id int) (aktorRelations []AktorRelation, err error) {

	//result := r.db.Where("fraaktørid = ?", id).Find(&aktorRelations)
	result := r.db.Table("AktørAktør").Select("AktørAktør.id, AktørAktør.fraaktørid, AktørAktør.tilaktørid, AktørAktørRolle.rolle AS rolle, Aktør.navn AS TilAktorNavn").Joins("left join AktørAktørRolle on AktørAktørRolle.id = AktørAktør.rolleid").Joins("left join Aktør on AktørAktør.tilaktørid = Aktør.id").Where("fraaktørid = ?", id).Find(&aktorRelations)
	err = result.Error
	fmt.Printf("%+v\n", aktorRelations)
	return 
}
