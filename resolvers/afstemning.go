package resolvers

import (
	"database/sql"
	"fmt"
	"time"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
)

type AfstemningResolver struct {
	afstemning Afstemning
}

func NewAfstemningList(args QueryArgs) (resolvers []*AfstemningResolver, err error){

	repo := newSqlite()

	query := "SELECT Afstemning.id, Afstemning.nummer, Afstemning.konklusion, Afstemning.vedtaget, Afstemning.kommentar, Afstemning.mødeid, Afstemningstype.type, Afstemning.sagstrinid, Afstemning.opdateringsdato FROM Afstemning JOIN Afstemningstype ON Afstemning.typeid = Afstemningstype.id"

	if args.Id != nil {
		query +=  " WHERE Afstemning.id=" + fmt.Sprintf("%d", *args.Id)
	}

	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	query+= " LIMIT " + os.Getenv("GRAPHQL_QUERY_LIMIT") + " OFFSET " + fmt.Sprintf("%d", *args.Offset) + ";"

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
	}

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
	return a.afstemning.Id
}

func (a *AfstemningResolver) Nummer() int32 {
	return a.afstemning.Nummer
}

func (a *AfstemningResolver) Konklusion() *string {
	return &a.afstemning.Konklusion
}

func (a *AfstemningResolver) Vedtaget() int32 {
	return a.afstemning.Vedtaget
}

func (a *AfstemningResolver) Kommentar() *string {
	return &a.afstemning.Kommentar
}

func (a *AfstemningResolver) Type() string {
	return a.afstemning.Type
}

func (a *AfstemningResolver) Opdateringsdato() graphql.Time {
	return a.afstemning.Opdateringsdato
}

func (a *AfstemningResolver) Møde()  (*MødeResolver, error) {
	margs := QueryArgs{Id:&a.afstemning.MødeID}
	return NewMøde(margs)
}

func (a *AfstemningResolver) Stemmer()  ([]*StemmeResolver, error) {
	// This introduces the N+1 problem
	

	args := StemmeQueryArgs{AfstemningId:&a.afstemning.Id}
	return NewStemmeList(args)
}