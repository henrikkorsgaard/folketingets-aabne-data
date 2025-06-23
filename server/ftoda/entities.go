package ftoda

type Sag struct {
	Id    int    `gorm:"primaryKey" json:"id"`
	Titel string `json:"titel"`
	//TitelKort         string `gorm:"column:titelkort" json:"titelkort"`
	Offentlighedskode string `gorm:"column:offentlighedskode"`
	//Nummer                 string
	//NummerPrefix           string `gorm:"column:nummerprefix"`
	//NummerNumerisk         string `gorm:"column:nummernumerisk"`
	//NummerPostfix          string `gorm:"column:nummerpostfix"`
	Resume                string `json:"resume"`
	Afstemningskonklusion string `gorm:"column:afstemningskonklusion" json:"afstemningskonklusion"`
	//PeriodeId              int
	//AfgorelsesResultatKode string `gorm:"column:afgorelsesresultatkode"`
	//Baggrundsmateriale     string
	//Opdateringsdato        string
	//StatsbudgetSag         int
	//Begrundelse         string
	//Paragrafnummer      int
	//Paragraf            string
	//AfgorelsesDato      string
	//Afgorelse           string
	//RådsmodeDato        string
	Lovnummer string `gorm:"column:lovnummer" json:"lovnummer"`
	//LovnummerDato       string
	//Retsinformationsurl string
	//FremsatUnderSagId   int
	//DeltUnderSagId      int

	// Foreign types
	//Type string `json:"sagstype"` //Table Sagtype
	//Kategori string `json:"sagkategori"` //Table Sagkategori
	//Status     string //Table Sagsstatus
	//SagstrinId int    //Table Sagstrin, use-case when identifying sag by sagstrin (relation Afstemning)
}
