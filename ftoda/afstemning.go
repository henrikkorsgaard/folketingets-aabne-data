package ftoda

//"gorm.io/gorm"
//graphql "github.com/graph-gophers/graphql-go"

type Afstemning struct {
	Id              int `gorm:"primaryKey"`
	Nummer          int
	Konklusion      string
	Vedtaget        int
	Kommentar       string
	MødeID          int
	Type            string
	SagstrinID      int
	Opdateringsdato string
}

func (Afstemning) TableName() string {
	return "Afstemning"
}

func (r *Repository) GetAfstemning(id int) (afstemning Afstemning, err error) {
	result := r.db.First(&afstemning, id)
	err = result.Error
	return
}

func (r *Repository) GetAllAfstemning(limit int, offset int) (afstemninger []Afstemning, err error) {
	result := r.db.Limit(limit).Offset(offset).Find(&afstemninger)
	err = result.Error

	return
}

func (r *Repository) GetAfstemningByIds(ids []int) (afstemninger []Afstemning, err error) {
	result := r.db.Find(&afstemninger, ids)
	err = result.Error

	return
}
