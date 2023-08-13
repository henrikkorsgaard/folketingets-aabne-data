package ftoda

type Stemme struct {
	Id              int `gorm:"primaryKey"`
	Type            string
	AfstemningId    int
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
