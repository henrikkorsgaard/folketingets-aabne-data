package resolvers

import (
	"henrikkorsgaard/folketingets-aabne-data/ftoda"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type StemmeQueryArgs struct {
	QueryArgs
	AfstemningId *int32
}

type StemmeResolver struct {
	stemme ftoda.Stemme
}

func NewStemmeList(args StemmeQueryArgs) (resolvers []*StemmeResolver, err error) {

	repo := ftoda.NewRepository()

	if args.Id != nil {
		var stemme ftoda.Stemme
		stemme, err = repo.GetStemme(int(*args.Id))

		stemmeResolver := StemmeResolver{stemme}
		resolvers = append(resolvers, &stemmeResolver)

		return
	}

	// if the query does not supply an offset
	// we set it to 0
	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	stemmer, err := repo.GetAllStemme(200, int(*args.Offset))

	for _, stemme := range stemmer {
		stemmeResolver := StemmeResolver{stemme}
		resolvers = append(resolvers, &stemmeResolver)
	}

	return
}

func NewStemme(args StemmeQueryArgs) (resolver *StemmeResolver, err error) {
	resolvers, err := NewStemmeList(args)
	if err != nil {
		return
	}

	if len(resolvers) > 0 {
		resolver = resolvers[0]
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

func (s *StemmeResolver) Aktør() (*AktørResolver, error) {
	id := int32(s.stemme.AktørId)
	args := AktørQueryArgs{QueryArgs: QueryArgs{Id: &id}}
	return NewAktør(args)
}

func (s *StemmeResolver) Afstemning() (*AfstemningResolver, error) {
	id := int32(s.stemme.AfstemningId)
	args := QueryArgs{Id: &id}
	return NewAfstemning(args)
}
