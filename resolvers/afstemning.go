package resolvers

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type AfstemningResolver struct {
	afstemning Afstemning
}

type Afstemning struct {
	Id int32
	Nummer int32
	Konklusion string
	Vedtaget int32
	Kommentar string
	MødeID int32 // Table: Møde , see how to resolve this with https://github.com/tonyghita/graphql-go-example/blob/37cd51aae44b998ee3baa2b7e9c21c56e11a5fe3/resolver/starship.go#L207
	Type string // Table: Afstemningstype 
	SagstrinID int32 // Table: Sagstrin
	Opdateringsdato graphql.Time //I have no idea what timezone, but given we do not operate on the update date, then accuracy is not important (famous last words)
}

func NewAfstemningList(args QueryArgs) (resolvers []*AfstemningResolver, err error){

	if args.Id != nil {
		afstemningResolver, err := NewAfstemning(args)
		if afstemningResolver != nil {
			resolvers = append(resolvers, afstemningResolver)
		}
		return resolvers, err
	}

	repo := newSqlite()

	query := "SELECT Afstemning.id, Afstemning.nummer, Afstemning.konklusion, Afstemning.vedtaget, Afstemning.kommentar, Afstemning.mødeid, Afstemningstype.type, Afstemning.sagstrinid, Afstemning.opdateringsdato FROM Afstemning JOIN Afstemningstype ON Afstemning.typeid = Afstemningstype.id;"

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

		afstemingResolver := AfstemningResolver{afstemning}
		resolvers = append(resolvers, &afstemingResolver)
	}

	if rows.Err() != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			err = fmt.Errorf("Unable to resolve Afstemning in database.")
		}
	}

	return
}

func NewAfstemning(args QueryArgs) (resolver *AfstemningResolver,err error) {
	repo := newSqlite()

	query := "SELECT Afstemning.id, Afstemning.nummer, Afstemning.konklusion, Afstemning.vedtaget, Afstemning.kommentar, Afstemning.mødeid, Afstemningstype.type, Afstemning.sagstrinid, Afstemning.opdateringsdato FROM Afstemning JOIN Afstemningstype ON Afstemning.typeid = Afstemningstype.id WHERE Afstemning.id=" + fmt.Sprintf("%d", *args.Id) + ";"
	
	afstemning := Afstemning{}
	var konklusion sql.NullString
	var kommentar sql.NullString
	var sagstringid sql.NullInt32
	var opdateringsdato string
	
	row := repo.db.QueryRow(query)
	
	err = row.Scan(&afstemning.Id, &afstemning.Nummer, &konklusion, &afstemning.Vedtaget, &kommentar, &afstemning.MødeID ,&afstemning.Type, &sagstringid, &opdateringsdato)

	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			err = fmt.Errorf("Unable to resolve Afstemning: Id %d does not exist", *args.Id)
		}
		return
	}

	t, err := time.Parse(time.DateTime, opdateringsdato)
	if err != nil {
		return
	}

	afstemning.Opdateringsdato.Time = t
	
	if konklusion.Valid {
		afstemning.Konklusion = konklusion.String
	}

	if kommentar.Valid {
		afstemning.Kommentar = kommentar.String
	}

	if sagstringid.Valid {
		afstemning.SagstrinID = sagstringid.Int32
	}

	resolver = &AfstemningResolver{afstemning}

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
	margs := QueryArgs{&a.afstemning.MødeID}
	return NewMøde(margs)
}