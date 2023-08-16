package resolvers

import (
	"fmt"
	"errors"
	"henrikkorsgaard/folketingets-aabne-data/ftoda"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type AktørQueryArgs struct {
	QueryArgs
	Type *string
	Navn *string
}

type AktørSearchArgs struct {
	QueryArgs
	Navn string
}

type AktørResolver struct {
	aktør ftoda.Aktør
}

func NewAktørList(args AktørQueryArgs) (resolvers []*AktørResolver, err error) {
	if args.Id != nil || args.Navn != nil {
		var aktørResolver *AktørResolver
		aktørResolver, err = NewAktør(args)
		resolvers = append(resolvers, aktørResolver)
		return
	}


	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	} 

	var aktører []ftoda.Aktør
	if args.Type != nil {
		aktører, err = ftoda.LoadAktørerByType(100, int(*args.Offset), *args.Type)
	} else {
		aktører, err = ftoda.LoadAktører(100, int(*args.Offset))
	}

	for _, aktør := range aktører {
		aktørResolver := AktørResolver{aktør}
		resolvers = append(resolvers, &aktørResolver)
	}

	return

}

func NewAktør(args AktørQueryArgs) (resolver *AktørResolver, err error) {

	// If there is an id and name, then we take id as the primary key and ignore name.
	// We could do LoadAktørByIdAndName 
	// It would return nil if there is an id / name conflict
	if args.Id != nil {
		id := int(*args.Id)
		var aktør ftoda.Aktør
		aktør, err = ftoda.LoadAktørById(id)
		if err != nil {
			err = errors.New("unable to resolve Aktør")
		}

		resolver = &AktørResolver{aktør}

		return
	}

	if args.Navn != nil {
		name := *args.Navn
		var aktør ftoda.Aktør
		aktør, err = ftoda.LoadAktørByName(name)
		if err != nil {
			err = errors.New("unable to resolve Aktør")
		}

		resolver = &AktørResolver{aktør}
		
		return 
	}

	err = errors.New("unable to resolve Aktør")
	
	return
} 

func NewAktørResultList(args AktørSearchArgs) (resolvers []*AktørResolver, err error) {
	
	aktører, err := ftoda.SearchAktørByName(100,args.Navn)
	fmt.Println(args.Navn)
	if err != nil {
		err = errors.New("unable to resolve Aktør search by name")
	}
	fmt.Println(len(aktører))
	for _, aktør := range aktører {
		aktørResolver := AktørResolver{aktør}
		resolvers = append(resolvers, &aktørResolver)
	}
	fmt.Println(len(resolvers))

	return
}

func (a *AktørResolver) Id() int32 {
	return int32(a.aktør.Id)
}

func (a *AktørResolver) Type() string {
	return a.aktør.Type
}

func (a *AktørResolver) GruppeNavnKort() *string {
	return &a.aktør.GruppeNavnKort
}

func (a *AktørResolver) Navn() *string {
	return &a.aktør.Navn
}

func (a *AktørResolver) Fornavn() *string {
	return &a.aktør.Fornavn
}

func (a *AktørResolver) Efternavn() *string {
	return &a.aktør.Efternavn
}

func (a *AktørResolver) Biografi() *string {
	return &a.aktør.Biografi
}

func (a *AktørResolver) Periode() *int32 {
	periode := int32(a.aktør.Periode)
	return &periode
}

func (a *AktørResolver) Startdato() *graphql.Time {
	return nil
}

func (a *AktørResolver) Slutdato() *graphql.Time {
	return nil
}

func (a *AktørResolver) Opdateringsdato() graphql.Time {
	t, err := time.Parse(time.DateTime, a.aktør.Opdateringsdato)
	if err != nil {
		panic(err) // This field is not null, I want to catch errors in development
	}
	return graphql.Time{t}
}

func (a *AktørResolver) Stemmer() (*[]*StemmeResolver, error) {
	id := int32(a.aktør.Id)
	args := StemmeQueryArgs{AktørId: &id}
	stemmeResolvers, err := NewStemmeList(args)
	return &stemmeResolvers, err
}


