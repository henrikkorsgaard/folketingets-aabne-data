package resolvers

import (

	graphql "github.com/graph-gophers/graphql-go"
)

type AktørQueryArgs struct {
	QueryArgs
	Type *string
}

type AktørResolver struct {
	aktør Aktør
}


func NewAktørList(args AktørQueryArgs) (resolvers []*AktørResolver, err error){
	/*
	
	if args.Type != nil {
		// query by type
	}

	if args.Id != nil {
		// query by type
	}

	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	

	*/

	return
}

//This will still be subject to N+1 unless I load it here. Other entities might call it from a list and then this will result in N+1 calls into NewAktør. Use AktørLoader
func NewAktør(args AktørQueryArgs) (resolver *AktørResolver,err error) {
	resolvers, err := NewAktørList(args)
	if err != nil {
		return
	}

	if len(resolvers) > 0 {
		resolver = resolvers[0]
	}

	return 
}

func (a *AktørResolver) Id() int32 {
	return a.aktør.Id
}

func (a *AktørResolver) Type() string {
	return a.aktør.Type
}

func (a *AktørResolver) Gruppenavnkort() *string {
	return &a.aktør.GruppeNavnKort
}

func (a *AktørResolver) Navn() *string {
	return &a.aktør.Navn
}

func (a *AktørResolver) Fornavn() *string {
	return &a.aktør.Navn
}

func (a *AktørResolver) Efternavn() *string {
	return &a.aktør.Navn
}

func (a *AktørResolver) Biografi() *string {
	return &a.aktør.Navn
}

//TODO: this should return a periodeResolver
func (a *AktørResolver) PeriodeID() *int32 {
	return &a.aktør.PeriodeID
}

func (a *AktørResolver) Startdato() *graphql.Time {
	return &a.aktør.StartDato
}

func (a *AktørResolver) Slutdato() *graphql.Time {
	return &a.aktør.SlutDato
}


func (a *AktørResolver) Opdateringsdato() graphql.Time {
	return a.aktør.Opdateringsdato
}
