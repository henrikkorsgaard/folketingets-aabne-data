package resolvers 

import (
	"errors"
	"henrikkorsgaard/folketingets-aabne-data/ftoda"

	"time"
	graphql "github.com/graph-gophers/graphql-go"
)

type SagQueryArgs struct {
	QueryArgs
	Type *string
}

type SagResolver struct {
	sag ftoda.Sag 
}

func NewSagList(args SagQueryArgs) (resolvers []*SagResolver, err error) {
	
	if args.Id != nil {
		var sagResolver *SagResolver
		sagResolver, err = NewSag(args)
		resolvers = append(resolvers, sagResolver)

		return
	}

	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	var sager []ftoda.Sag 

	if args.Type != nil {
		sager, err = ftoda.LoadSagerByType(100, int(*args.Offset), *args.Type)
	} else {
		sager, err = ftoda.LoadSager(100, int(*args.Offset))
	}

	for _, s := range sager {
		sagResolver := SagResolver{s}
		resolvers = append(resolvers, &sagResolver)
	}

	return
}

func NewSag(args SagQueryArgs) (resolver *SagResolver, err error) {
	if args.Id != nil {
		id := int(*args.Id)

		var sag ftoda.Sag 
		sag, err = ftoda.LoadSag(id)
		if err != nil {
			err = errors.New("Unable to resolve Sag: " + err.Error())
		}

		resolver = &SagResolver{sag}
		return
	}

	err = errors.New("unable to resolve Sag")
	return
}

func (s *SagResolver)Id() int32 {
	return int32(s.sag.Id)
}

func (s *SagResolver)Titel() *string {
	return &s.sag.Titel 
}

func (s *SagResolver)Titelkort() *string {
	return &s.sag.TitelKort
}

func (s *SagResolver)Resume() *string {
	return &s.sag.Resume
}

func (s *SagResolver)Afstemningskonklusion() *string {
	return &s.sag.Afstemningskonklusion
}

func (s *SagResolver)Baggrundsmateriale() *string {
	return &s.sag.Baggrundsmateriale
}

func (s *SagResolver)Begrundelse() *string {
	return &s.sag.Begrundelse
}

func (s *SagResolver)Lovnummer() *string {
	return &s.sag.Lovnummer
}

func (s *SagResolver)Paragrafnummer() *int32 {
	pnum := int32(s.sag.Paragrafnummer)
	return &pnum
}

func (s *SagResolver)Paragraf() *string {
	return &s.sag.Paragraf
}

func (s *SagResolver)Retsinformationsurl() *string {
	return &s.sag.Retsinformationsurl
}

func (s *SagResolver) Opdateringsdato() graphql.Time {
	t, err := time.Parse(time.DateTime, s.sag.Opdateringsdato)
	if err != nil {
		panic(err) // This field is not null, I want to catch errors in development
	}
	return graphql.Time{t}
}

func (s *SagResolver)Type() *string {
	return &s.sag.Type
}

func (s *SagResolver)Kategori() *string {
	return &s.sag.Kategori
}

func (s *SagResolver)Status() *string {
	return &s.sag.Status 
}