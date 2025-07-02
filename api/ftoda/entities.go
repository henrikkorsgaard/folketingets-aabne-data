package ftoda

type Sag struct {
	Id                int    `gorm:"primaryKey" json:"id"`
	Titel             string `json:"titel"`
	TitelKort         string `gorm:"column:titelkort" json:"titelkort"`
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
}

type Afstemning struct {
	Id              int `gorm:"primaryKey" json:"id"`
	Nummer          int
	Konklusion      string `gorm:"column:konklusion" json:"konklusion"`
	Vedtaget        bool   `gorm:"column:vedtaget" json:"vedtaget"`
	Kommentar       string `gorm:"column:kommentar" json:"kommentar"`
	ModeId          int    `gorm:"column:mødeid"`
	Type            string `gorm:"column:type" json:"type"`
	SagstrinId      int
	SagId           int `gorm:"column:sagid" json:"sagid"`
	Opdateringsdato string
}

// This is the relationship between Stemme and actor
type Stemme struct {
	Id              int    `gorm:"primaryKey" json:"id"`
	Type            string `gorm:"type" json:"type"`
	AfstemningId    int    `gorm:"column:afstemningid"`
	AktorId         int    `gorm:"column:aktørid"`
	Opdateringsdato string
}

type Aktor struct {
	Id             int    `gorm:"primaryKey" json:"id"`
	Type           string `gorm:"type" json:"type"`
	GruppeNavnKort string `gorm:"column:gruppenavnkort" json:"gruppenavnkort"`
	Navn           string `gorm:"navn" json:"navn"`
	//Fornavn         string
	//Efternavn       string
	//Biografi        string
	//Periode         int
	//Opdateringsdato string
	//Startdato       string
	//Slutdato        string
}

type Sagstrin struct {
	Id              int    `gorm:"primaryKey" json:"id"`
	Titel           string `gorm:"titel" json:"titel"`
	Sagid           int    `gorm:"sagid" json:"sagid"`
	Type            string `gorm:"type" json:"type"`
	Typeid          int    `gorm:"column:typeid" json:"typeid"`
	Statusid        int    `gorm:"column:statusid" json:"statusid"`
	Dato            string
	Opdateringsdato string
}
