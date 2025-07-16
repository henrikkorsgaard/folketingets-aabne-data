package ftoda

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

/*
	Note:
	The ftoda data contains 2 date variants
	dato: "2025-08-14T00:00:00"
	opdateringsdato: "2025-06-26T12:03:29.443"

	For now, we only need the first (dato). The second is likely a database related date.

	See: https://oda.ft.dk/api/Sagstrin?$format=json&$filter=sagid%20eq%20102467&$skip=0&orderby=id%20desc
*/

type FtodaDate struct {
	time.Time //make a gorm data alias instead?
}

func (d *FtodaDate) Value() (driver.Value, error) {
	return driver.Value(d.Time.Format("2006-01-02T15:04:05")), nil
}

func (d *FtodaDate) Scan(value interface{}) error {
	// TODO: Write scanner for gorm https://stackoverflow.com/a/65459041
	fmt.Println(value)
	return nil
}

func (d *FtodaDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

type Sag struct {
	Id                    int    `gorm:"primaryKey" json:"id"`
	Titel                 string `gorm:"column:titel" json:"titel"`
	TitelKort             string `gorm:"column:titelkort" json:"titelkort"`
	Nummer                string `gorm:"column:nummer" json:"nummer"` //e.g. L 105
	Resume                string `json:"resume"`
	Afstemningskonklusion string `gorm:"column:afstemningskonklusion" json:"afstemningskonklusion"`
	StatusId              int    `gorm:"column:statusid" json:"statusid"`
	//PeriodeId             int    //maps onto
	Begrundelse    string `gorm:"column:begrundelse" json:"begrundelse"`
	Paragrafnummer int    `gorm:"column:paragrafnummer" json:"paragrafnummer"`
	Paragraf       string `gorm:"column:paragraf" json:"paragraf"`
	Lovnummer      string `gorm:"column:lovnummer" json:"lovnummer"`

	//Opdateringsdato datatypes.Date
}

// Helper type for analysis
// TODO: Make relational for database optimization
type SagSagstrin struct {
	Sag
	Sagstrin []SagSagstrin `json:"sagstrin"`
}

type EmneordSag struct {
	Id        int `gorm:"primarykey" json:"id"`
	EmneordId int `gorm:"emneordid" json:"emneordid"`
	SagId     int `gorm:"sagid" json:"sagid"`
	//Opdateringsdato datatypes.Date
}

type Emneord struct {
	Id         int        `gorm:"primarykey" json:"id"`
	Emneord    string     `gorm:"emneord" json:"emneord"`
	TypeId     int        `gorm:"typeid" json:"typeid"`
	EmneordSag EmneordSag `gorm:"-" json:"-"`
	//Opdateringsdato datatypes.Date
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

type Stemme struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	Type         string `gorm:"type" json:"type"`
	AfstemningId int    `gorm:"column:afstemningid"`
	AktorId      int    `gorm:"column:aktørid"`
	//Opdateringsdato datatypes.Date
}

type Aktor struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	Type           string    `gorm:"type" json:"type"`
	GruppeNavnKort string    `gorm:"column:gruppenavnkort" json:"gruppenavnkort"`
	Navn           string    `gorm:"navn" json:"navn"`
	Fornavn        string    `gorm:"fornavn" json:"fornavn"`
	Efternavn      string    `gorm:"efternavn" json:"efternavn"`
	Biografi       string    `gorm:"biografi" json:"biografi"`
	Periode        int       `gorm:"periode" json:"periode"`
	Startdato      FtodaDate `gorm:"startdato" json:"startdato"`
	Slutdato       FtodaDate `gorm:"slutdato" json:"slutdato"`
	//Opdateringsdato datatypes.Date
}

type Sagstrin struct {
	Id       int       `gorm:"primaryKey" json:"id"`
	Titel    string    `gorm:"titel" json:"titel"`
	Sagid    int       `gorm:"sagid" json:"sagid"`
	Typeid   int       `gorm:"column:typeid" json:"typeid"`
	Statusid int       `gorm:"column:statusid" json:"statusid"`
	Dato     FtodaDate `gorm:"column:dato" json:"dato"`
	//Opdateringsdato datatypes.Date
}

type Sagstrinstype struct {
	Id   int       `gorm:"primaryKey" json:"id"`
	Type string    `gorm:"type" json:"type"`
	Dato FtodaDate `gorm:"column:dato" json:"dato"`
	//Opdateringsdato datatypes.Date
}

type SagstrinAktør struct {
	Id         int `gorm:"primaryKey" json:"id"`
	SagstrinId int `gorm:"sagstrinid" json:"sagstrinid"`
	AktørId    int `gorm:"aktørid" json:"aktørid"`
	RolleId    int `gorm:"rolleid" json:"rolleid"`
	//Opdateringsdato datatypes.Date
}
