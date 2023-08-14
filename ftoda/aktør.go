package ftoda



type Aktør struct {
	Id int32 
	Type string
	GruppeNavnKort string
	Navn string
	Fornavn string
	Efternavn string
	Biografi string 
	PeriodeId int32
	Stemmer []Stemme
	StartDato graphql.Time
	SlutDato graphql.Time
	Opdateringsdato graphql.Time
}