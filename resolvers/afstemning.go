package resolvers

import (

	"time"

	"henrikkorsgaard/folketingets-aabne-data/ftoda"

	graphql "github.com/graph-gophers/graphql-go"
)

type AfstemningResolver struct {
	afstemning ftoda.Afstemning
}

func NewAfstemningList(args QueryArgs) (resolvers []*AfstemningResolver, err error){

	repo := ftoda.NewRepository()

	if args.Id != nil {
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

	afstemninger, err := repo.GetAllAfstemning(100, int(*args.Offset))

	for _, afstemning := range afstemninger {
		afstemingResolver := AfstemningResolver{afstemning}
		resolvers = append(resolvers, &afstemingResolver)
	}
	/*
	rows, err := repo.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next(){
		afstemning := Afstemning{}
		var konklusion sql.NullString
		var kommentar sql.NullString
		var sagstringid sql.NullInt32
		var opdateringsdato string

		err = rows.Scan(&afstemning.Id, &afstemning.Nummer, &konklusion, &afstemning.Vedtaget, &kommentar, &afstemning.MødeID ,&afstemning.Type, &sagstringid, &opdateringsdato)

		if err != nil {
			break
		}

		t_op, err := time.Parse(time.DateTime, opdateringsdato)
		if err != nil {
			break
		}

		afstemning.Opdateringsdato.Time = t_op

		afstemingResolver := AfstemningResolver{afstemning}
		resolvers = append(resolvers, &afstemingResolver)
	}

	if rows.Err() != nil {
		err = rows.Err()
	}

	if args.Id != nil && len(resolvers) == 0 {
		err = fmt.Errorf("Unable to resolve Afstemning: Id %d does not exist", *args.Id)
	}*/

	return
}

func NewAfstemning(args QueryArgs) (resolver *AfstemningResolver,err error) {

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

func (a *AfstemningResolver) Møde()  (*MødeResolver, error) {
	id := int32(a.afstemning.MødeID)
	margs := QueryArgs{Id:&id}
	return NewMøde(margs)
}

func (a *AfstemningResolver) Stemmer()  ([]*StemmeResolver, error) {
	// This introduces the N+1 problem
	id := int32(a.afstemning.Id)
	args := StemmeQueryArgs{AfstemningId:&id}
	return NewStemmeList(args)
}