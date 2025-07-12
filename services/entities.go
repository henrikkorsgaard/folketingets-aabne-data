package ftoda

import (
	"strings"
	"time"
)

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

type FtodaDate struct {
	time.Time //make a gorm data alias instead?
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
	Id                    int          `gorm:"primaryKey" json:"id"`
	Titel                 string       `json:"titel"`
	TitelKort             string       `gorm:"column:titelkort" json:"titelkort"`
	Nummer                string       `gorm:"column:nummer" json:"nummer"` //e.g. L 105
	Resume                string       `json:"resume"`
	Afstemningskonklusion string       `gorm:"column:afstemningskonklusion" json:"afstemningskonklusion"`
	PeriodeId             int          //maps onto
	Begrundelse           string       `gorm:"column:begrundelse" json:"begrundelse"`
	Paragrafnummer        int          `gorm:"column:paragrafnummer" json:"paragrafnummer"`
	Paragraf              string       `gorm:"column:paragraf" json:"paragraf"`
	Lovnummer             string       `gorm:"column:lovnummer" json:"lovnummer"`
	EmneordSager          []EmneordSag `gorm:"emneordsag" json:"emneordsag"`
	//Opdateringsdato datatypes.Date

	/* Relational values extended without persisting this */
	Emneord []string `gorm:"-"`
}

type EmneordSag struct {
	Id        int `gorm:"primarykey" json:"id"`
	EmneordId int `gorm:"emneordid" json:"emneordid"`
	SagId     int `gorm:"sagid" json:"sagid"`
}

type Emneord struct {
	Id      int    `gorm:"primarykey" json:"id"`
	Emneord string `gorm:"emneord" json:"emneord"`
	TypeId  int    `gorm:"typeid" json:"typeid"`
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

/*
The official representation of legislation sort the sagstrin into buckets:

- Fremsættelse: When the legislation is proposed
	- Fremsættelse typeid: 6,
	- Lovforslag som fremsat
- 1. Behandling: The initial debate of the legislation
	- 1. Behandling
	-
- 2. Behandling: The second debate, change and vote
- 3. Behandling: The final debate, change and vote
- Udvalgsbehandling: Between the step 1. and 2. and 2. and 3., legislation can go into domain/area specific committee (e.g. tax)

We need to just bucket things and then handle if they fall outside

*/

type Sagstrin struct {
	Id            int           `gorm:"primaryKey" json:"id"`
	Titel         string        `gorm:"titel" json:"titel"`
	Sagid         int           `gorm:"sagid" json:"sagid"`
	Typeid        int           `gorm:"column:typeid" json:"typeid"`
	Statusid      int           `gorm:"column:statusid" json:"statusid"`
	Sagstrinstype Sagstrinstype `gorm:"sagstrinstype" json:"sagstrinstype"`
	Dato          FtodaDate     `gorm:"column:dato" json:"dato"`
	Afstemning    []Afstemning  `gorm:"column:afstemning" json:"afstemning"`
	//Opdateringsdato datatypes.Date
}

type Sagstrinstype struct {
	Id   int       `gorm:"primaryKey" json:"id"`
	Type string    `gorm:"type" json:"type"` //this should be a string
	Dato FtodaDate `gorm:"column:dato" json:"dato"`
	//Opdateringsdato datatypes.Date
}

type LovforslagsType int

// see example here: https://www.ft.dk/samling/20131/lovforslag/l160/index.htm

const (
	Fremsættelse LovforslagsType = iota
	FørsteBehandling
	AndenBehandling
	TredjeBehandling
	FørsteUdvalgsbehandling
	AndenUdvalgsbehandling
	NoClue //for catching stuff we want to do something about
)

func (lft LovforslagsType) String() string {
	return [...]string{"Fremsættelse", "1. Behandling", "2. Behandling", "3. Behandling", "Udvalgsbehandling", "2. Udvalgsbehandling", "NoClue"}[lft]
}

type LovforslagsTrin struct {
	//this is the parent
	Id       int        `gorm:"primaryKey" json:"id"`
	Sagstrin []Sagstrin `gorm:"sagstrin" json:"sagstrin"`
	Sagid    int        `gorm:"sagid" json:"sagid"`
	Type     string     `gorm:"type" json:"type"`
}

func InitiateLovforslagsTrin(sagid int) []LovforslagsTrin {
	m := make([]LovforslagsTrin, 7)

	m[0] = LovforslagsTrin{Sagid: sagid, Type: LovforslagsType(0).String()}
	m[1] = LovforslagsTrin{Sagid: sagid, Type: LovforslagsType(1).String()}
	m[2] = LovforslagsTrin{Sagid: sagid, Type: LovforslagsType(2).String()}
	m[3] = LovforslagsTrin{Sagid: sagid, Type: LovforslagsType(3).String()}
	m[4] = LovforslagsTrin{Sagid: sagid, Type: LovforslagsType(4).String()}
	m[5] = LovforslagsTrin{Sagid: sagid, Type: LovforslagsType(5).String()}
	m[6] = LovforslagsTrin{Sagid: sagid, Type: LovforslagsType(6).String()}

	return m
}

type SagstrinAktør struct {
	Id         int `gorm:"primaryKey" json:"id"`
	SagstrinId int `gorm:"sagstrinid" json:"sagstrinid"`
	AktørId    int `gorm:"aktørid" json:"aktørid"`
	RolleId    int `gorm:"rolleid" json:"rolleid"`
	//Opdateringsdato datatypes.Date
}
