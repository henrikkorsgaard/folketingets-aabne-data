package resolvers

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"henrikkorsgaard/folketingets-aabne-data/ftoda"
)

type StemmeQueryArgs struct {
	QueryArgs
	AfstemningId *int32
	AktorId *int32
}

type StemmeResolver struct {
	stemme ftoda.Stemme
}

func NewStemmeList(args StemmeQueryArgs) (resolvers []*StemmeResolver, err error) {

	if args.AfstemningId != nil {
		id := int(*args.AfstemningId)
		var stemmer []ftoda.Stemme
		stemmer, err = ftoda.LoadStemmerFromAfstemning(id)

		for _, stemme := range stemmer {
			stemmeResolver := StemmeResolver{stemme}
			resolvers = append(resolvers, &stemmeResolver)
		}
	}

	if args.AktorId != nil {
		id := int(*args.AktorId)
		var stemmer []ftoda.Stemme
		stemmer, err = ftoda.LoadStemmerFromAktor(id)

		for _, stemme := range stemmer {
			stemmeResolver := StemmeResolver{stemme}
			resolvers = append(resolvers, &stemmeResolver)
		}
	}

	return
}

func (s *StemmeResolver) Id() int32 {
	return int32(s.stemme.Id)
}

func (s *StemmeResolver) Type() *string {
	return &s.stemme.Type
}


func (s *StemmeResolver) Opdateringsdato() graphql.Time {
	t, err := time.Parse(time.DateTime, s.stemme.Opdateringsdato)
	if err != nil {
		panic(err) // This field is not null, I want to catch errors in development
	}
	return graphql.Time{t}
}

func (s *StemmeResolver) Aktor() (*AktorResolver, error) {
	id := int32(s.stemme.AktorId)
	args := AktorQueryArgs{QueryArgs: QueryArgs{Id: &id}}
	return NewAktor(args)
}

func (s *StemmeResolver) Afstemning() (*AfstemningResolver, error) {
	id := int32(s.stemme.AfstemningId)
	args := AfstemningQueryArgs{QueryArgs:QueryArgs{Id: &id}}
	return NewAfstemning(args)
}
