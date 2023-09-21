package resolvers

import (
	"errors"
	"henrikkorsgaard/folketingets-aabne-data/ftoda"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type AfstemningQueryArgs struct {
	QueryArgs
	Type *string
	Kommentar *bool
	SagId *int32
}

type AfstemningResolver struct {
	afstemning ftoda.Afstemning
}

// This is the pattern to follow.
func NewAfstemningList(args AfstemningQueryArgs) (resolvers []*AfstemningResolver, err error) {

	if args.Id != nil {
		var afstemningResolver *AfstemningResolver
		afstemningResolver, err = NewAfstemning(args)
		resolvers = append(resolvers, afstemningResolver)
		
		return
	}

	// if the query does not supply an offset
	// we set it to 0
	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	var afstemninger []ftoda.Afstemning

	if args.Type != nil {
		afstemninger, err = ftoda.LoadAfstemningerByType(100, int(*args.Offset), *args.Type) 
	} else if args.Kommentar != nil && *args.Kommentar {
		afstemninger, err = ftoda.LoadAfstemningerWithKommentar(100, int(*args.Offset)) 
	} else if args.SagId != nil {
		afstemninger, err = ftoda.LoadAfstemningerBySag(int(*args.SagId))
	} else {
		//load all i guess
		afstemninger, err = ftoda.LoadAfstemninger(100, int(*args.Offset)) 
	}

	for _, afstemning := range afstemninger {
		afstemingResolver := AfstemningResolver{afstemning}
		resolvers = append(resolvers, &afstemingResolver)
	}

	return
}

func NewAfstemning(args AfstemningQueryArgs) (resolver *AfstemningResolver, err error) {
	
	if args.Id != nil {
		
		id := int(*args.Id)
		
		var afstemning ftoda.Afstemning
	
		afstemning, err = ftoda.LoadAfstemning(id)

		if err != nil {
			err = errors.New("Unable to resolve Afstemning: " + err.Error())
		}
		resolver = &AfstemningResolver{afstemning}	
		return
	}
	err = errors.New("unable to resolve Afstemning")

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

func (a *AfstemningResolver) Mode() int32 {
	return int32(a.afstemning.ModeId)
}

func (a *AfstemningResolver) Stemmer() ([]*StemmeResolver, error) {
	id := int32(a.afstemning.Id)
	args := StemmeQueryArgs{AfstemningId: &id}
	return NewStemmeList(args)
}

func (a *AfstemningResolver) Sag() (*SagResolver, error) {
	id := int32(a.afstemning.SagId)
	args := SagQueryArgs{QueryArgs: QueryArgs{Id:&id}}
	return NewSag(args)
}
