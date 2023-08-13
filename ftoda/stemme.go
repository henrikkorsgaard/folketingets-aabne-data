package ftoda

type Stemme struct {
	Id int 					`gorm:"primaryKey"`
	Type string
	AfstemningId int
	AktørID int
	Opdateringsdato string
}