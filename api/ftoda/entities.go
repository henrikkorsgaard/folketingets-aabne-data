package ftoda

import (
	"fmt"
	"strings"
	"time"
)

type FtodaDate struct {
	time.Time //make a gorm data alias instead?
}

// See https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json

/*
	Note:
	The ftoda data contains 2 date variants
	dato: "2025-08-14T00:00:00"
	opdateringsdato: "2025-06-26T12:03:29.443"

	For now, we only need the first (dato). The second is likely a database related date.

	See: https://oda.ft.dk/api/Sagstrin?$format=json&$filter=sagid%20eq%20102467&$skip=0&orderby=id%20desc
*/

/*
	see: https://go.dev/play/p/14Un1FWtUf6

	We trim the datetime to only have date. Because that is all we need.

*/

func (d *FtodaDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
	/*
		switch len(s) {
		case 10:
			theTime, err = time.Parse("2006-01-02", s)
		case 16:
			theTime, err = time.Parse("2006-01-02T15:04", s)
		case 19:
			theTime, err = time.Parse("2006-01-02T15:04:05", s)
		default:
			err = json.Unmarshal(b, &theTime)
			if err != nil {
				err = fmt.Errorf("layout must: 2006-01-02T15:04:05-07:00 | %w", err)
			}
		}*/

}

func (d FtodaDate) MarshalJSON() ([]byte, error) {
	fmt.Println("called")
	t := fmt.Sprintf("\"%s\"", time.Time(d.Time).Format("2006-01-02"))
	return []byte(t), nil
}

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
	Id         int `gorm:"primaryKey" json:"id"`
	Nummer     int
	Konklusion string `gorm:"column:konklusion" json:"konklusion"`
	Vedtaget   bool   `gorm:"column:vedtaget" json:"vedtaget"`
	Kommentar  string `gorm:"column:kommentar" json:"kommentar"`
	ModeId     int    `gorm:"column:mødeid"`
	Type       string `gorm:"column:type" json:"type"`
	SagstrinId int
	SagId      int `gorm:"column:sagid" json:"sagid"`
	//Opdateringsdato datatypes.Date
}

// This is the relationship between Stemme and actor
type Stemme struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	Type         string `gorm:"type" json:"type"`
	AfstemningId int    `gorm:"column:afstemningid"`
	AktorId      int    `gorm:"column:aktørid"`
	//Opdateringsdato datatypes.Date
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
	Id       int       `gorm:"primaryKey" json:"id"`
	Titel    string    `gorm:"titel" json:"titel"`
	Sagid    int       `gorm:"sagid" json:"sagid"`
	Type     string    `gorm:"type" json:"type"`
	Typeid   int       `gorm:"column:typeid" json:"typeid"`
	Statusid int       `gorm:"column:statusid" json:"statusid"`
	Dato     FtodaDate `gorm:"column:dato" json:"dato"`
	//Opdateringsdato datatypes.Date
}
