package ftoda 


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
	Afstemingskonklusion string `gorm:"column:afstemningskonklusion"`
	PeriodeId int 
	AfgørelsesResultatKode string `gorm:"column:afgørelsesresultatkode"`
	Baggrundsmateriale string 
	Opdateringsdato string
	StatsbudgetSag int 
	Begrundelse string
	Paragrafnummer int
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

func (r *Repository) getSager(limit, offset int) (sager []Sag, err error){
	result := r.db.Table("Sag").Limit(limit).Offset(offset).Select("Sag.*, Sagstype.type, Sagskategori.kategori, Sagsstatus.status").Joins("left join Sagstype on Sag.typeid = Sagstype.id").Joins("left join Sagskategori on Sag.kategoriid = Sagskategori.id").Joins("left join Sagsstatus on Sag.statusid = Sagsstatus.id").Find(&sager)
	err = result.Error
	return
}

