package resolvers

import (
	"context"
	"henrikkorsgaard/folketingets-aabne-data/ftoda"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type AfstemningResolver struct {
	afstemning ftoda.Afstemning
}

// This is the pattern to follow.
func NewAfstemningList(args QueryArgs) (resolvers []*AfstemningResolver, err error) {

	if args.Id != nil {
		// load one I guess
		var afstemning ftoda.Afstemning
		afstemning, err = repo.GetAfstemning(int(*args.Id))

		afstemingResolver := AfstemningResolver{afstemning}
		resolvers = append(resolvers, &afstemingResolver)
		return
	}

	// if the query does not supply an offset
	// we set it to 0
	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	//load all i guess
	afstemninger, err := repo.GetAllAfstemning(100, int(*args.Offset))

	for _, afstemning := range afstemninger {
		afstemingResolver := AfstemningResolver{afstemning}
		resolvers = append(resolvers, &afstemingResolver)
	}

	return
}

func NewAfstemning(args QueryArgs) (resolver *AfstemningResolver, err error) {

	resolvers, err := NewAfstemningList(args)
	if err != nil {
		return
	}

	if len(resolvers) > 0 {
		resolver = resolvers[0]
	}

	return
}

func (a *AfstemningResolver) Id() int32 {
	return int32(a.afstemning.Id)
}

func (a *AfstemningResolver) Nummer() int32 {
	return int32(a.afstemning.Nummer)
}

func (a *AfstemningResolver) Konklusion() *string {
	return &a.afstemning.Konklusion
}

func (a *AfstemningResolver) Vedtaget() int32 {
	return int32(a.afstemning.Vedtaget)
}

func (a *AfstemningResolver) Kommentar() *string {
	return &a.afstemning.Kommentar
}

func (a *AfstemningResolver) Type() string {
	return a.afstemning.Type
}

func (a *AfstemningResolver) Opdateringsdato() graphql.Time {
	t, err := time.Parse(time.DateTime, a.afstemning.Opdateringsdato)
	if err != nil {
		panic(err) // This field is not null, I want to catch errors in development
	}
	return graphql.Time{t}
}

func (a *AfstemningResolver) Møde() (*MødeResolver, error) {
	id := int32(a.afstemning.MødeID)
	margs := QueryArgs{Id: &id}
	return NewMøde(margs)
}

func (a *AfstemningResolver) Stemmer() ([]*StemmeResolver, error) {
	id := int32(a.afstemning.Id)
	args := StemmeQueryArgs{AfstemningId: &id}
	return NewStemmeList(context.Background(), args)
}
