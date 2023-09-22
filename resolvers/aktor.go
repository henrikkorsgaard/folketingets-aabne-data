package resolvers

import (
	"errors"
	"henrikkorsgaard/folketingets-aabne-data/ftoda"
	"time"
	
	graphql "github.com/graph-gophers/graphql-go"
)

type AktorQueryArgs struct {
	QueryArgs
	Type *string
	Navn *string
}

type AktorSearchArgs struct {
	QueryArgs
	Navn string
}

type AktorResolver struct {
	aktor ftoda.Aktor
}

func NewAktorList(args AktorQueryArgs) (resolvers []*AktorResolver, err error) {
	if args.Id != nil || args.Navn != nil {
		var aktorResolver *AktorResolver
		aktorResolver, err = NewAktor(args)
		resolvers = append(resolvers, aktorResolver)
		return
	}


	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	} 

	var aktorer []ftoda.Aktor
	if args.Type != nil {
		aktorer, err = ftoda.LoadAktorerByType(100, int(*args.Offset), *args.Type)
	} else {
		aktorer, err = ftoda.LoadAktorer(100, int(*args.Offset))
	}

	for _, aktor := range aktorer {
		aktorResolver := AktorResolver{aktor}
		resolvers = append(resolvers, &aktorResolver)
	}

	return

}

func NewAktor(args AktorQueryArgs) (resolver *AktorResolver, err error) {

	// If there is an id and name, then we take id as the primary key and ignore name.
	// We could do LoadAktorByIdAndName 
	// It would return nil if there is an id / name conflict
	if args.Id != nil {
		id := int(*args.Id)
		var aktor ftoda.Aktor
		aktor, err = ftoda.LoadAktorById(id)
		if err != nil {
			err = errors.New("unable to resolve Aktor")
		}

		resolver = &AktorResolver{aktor}

		return
	}

	if args.Navn != nil {
		name := *args.Navn
		var aktor ftoda.Aktor
		aktor, err = ftoda.LoadAktorByName(name)
		if err != nil {
			err = errors.New("unable to resolve Aktor")
		}

		resolver = &AktorResolver{aktor}
		
		return 
	}

	err = errors.New("unable to resolve Aktor")
	
	return
} 

func NewAktorResultList(args AktorSearchArgs) (resolvers []*AktorResolver, err error) {
	
	aktorer, err := ftoda.SearchAktorByName(100,args.Navn)
	if err != nil {
		err = errors.New("unable to resolve Aktor search by name")
	}

	for _, aktor := range aktorer {
		aktorResolver := AktorResolver{aktor}
		resolvers = append(resolvers, &aktorResolver)
	}

	return
}

func (a *AktorResolver) Id() int32 {
	return int32(a.aktor.Id)
}

func (a *AktorResolver) Type() string {
	return a.aktor.Type
}

func (a *AktorResolver) GruppeNavnKort() *string {
	return &a.aktor.GruppeNavnKort
}

func (a *AktorResolver) Navn() *string {
	return &a.aktor.Navn
}

func (a *AktorResolver) Fornavn() *string {
	return &a.aktor.Fornavn
}

func (a *AktorResolver) Efternavn() *string {
	return &a.aktor.Efternavn
}

func (a *AktorResolver) Biografi() *string {
	return &a.aktor.Biografi
}

func (a *AktorResolver) Periode() *int32 {
	periode := int32(a.aktor.Periode)
	return &periode
}

func (a *AktorResolver) Startdato() *graphql.Time {
	return nil
}

func (a *AktorResolver) Slutdato() *graphql.Time {
	return nil
}

func (a *AktorResolver) Opdateringsdato() graphql.Time {
	t, err := time.Parse(time.DateTime, a.aktor.Opdateringsdato)
	if err != nil {
		panic(err) // This field is not null, I want to catch errors in development
	}
	return graphql.Time{t}
}


// Comined or extracted fields 
func (a *AktorResolver) Stemmer() (*[]*StemmeResolver, error) {
	id := int32(a.aktor.Id)
	args := StemmeQueryArgs{AktorId: &id}
	stemmeResolvers, err := NewStemmeList(args)
	return &stemmeResolvers, err
}

func (a *AktorResolver) Fodselsdato() *graphql.Time {
	t, err := time.Parse("02-01-2006", a.aktor.Fodselsdato)
	if err != nil {
		return nil 
	}
	return &graphql.Time{t}
}

func (a *AktorResolver) Dodsdato() *graphql.Time {
	t, err := time.Parse("02-01-2006", a.aktor.Dodsdato)
	if err != nil {
		return nil 
	}
	return &graphql.Time{t}
}

func (a *AktorResolver) Personligt() *string {
	return &a.aktor.Personligt
}

func (a *AktorResolver) Parti() *string {
	return &a.aktor.Parti
}

func (a *AktorResolver) Billede() *string {
	return &a.aktor.Billede
}

func (a *AktorResolver) Kon() *string {
	return &a.aktor.Kon
}

func (a *AktorResolver) Uddannelsesniveau() *string {
	return &a.aktor.Uddannelsesniveau
}

func (a *AktorResolver) Titel() *string {
	return &a.aktor.Titel
}

func (a *AktorResolver) Uddannelse() *[]*string {
	var ref []*string 
	for i := range a.aktor.Uddannelse {
		ref = append(ref, &a.aktor.Uddannelse[i])
	}
	return &ref
}

func (a *AktorResolver) Beskaftigelse() *[]*string {
	var ref []*string 
	for i := range a.aktor.Beskaftigelse {
		ref = append(ref, &a.aktor.Beskaftigelse[i])
	}
	return &ref
}

func (a *AktorResolver) Ministerposter() *[]*string {
	var ref []*string 
	for i := range a.aktor.Ministerposter {
		ref = append(ref, &a.aktor.Ministerposter[i])
	}
	return &ref
}

func (a *AktorResolver) Valgkredse() *[]*string {
	var ref []*string 
	for i := range a.aktor.Valgkredse {
		ref = append(ref, &a.aktor.Valgkredse[i])
	}
	return &ref
}


