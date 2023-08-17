package ftoda 


type Sag struct {
	Id int `gorm:"primaryKey"`
	
	// Foreign types
	Type string //Table Sagtype 
	Kategori string //Table Sagkategori
	Status string  //Table Sagsstatus
}

func (Sag) TableName() string {
	return "Aktør"
}

func LoadSager(limit, offset int) (sager []Sag, err error) {
	repo := newRepository()
	return repo.getSager(limit, offset)
}

func (r *Repository) getSager(limit, offset int) (sager []Sag, err error){
	result := r.db.Table("Sag").Limit(limit).Offset(offset).Select("Sag.*, Sagstype.type, Sagskategori.kategori, Sagsstatus.status").Joins("left join Sagstype on Sag.typeid = Sagstype.id").Joins("left join Sagskategori on Sag.kategoriid = Sagskategori.id").Joins("left join Sagsstatus on Sag.statusid = Sagsstatus.id").Find(&sager)
	err = result.Error
	return
}

