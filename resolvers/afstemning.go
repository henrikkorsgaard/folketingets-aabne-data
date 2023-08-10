package resolvers

import (
	"database/sql"
	"fmt"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type AfstemningQueryArgs struct {
	Id *int 
}

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


func NewAfstemning(args AfstemningQueryArgs) (resolver *AfstemningResolver,err error) {
	
	repo := newSqlite()

	query := "SELECT Afstemning.id, Afstemning.nummer, Afstemning.konklusion, Afstemning.vedtaget, Afstemning.kommentar, Afstemning.mødeid, Afstemningstype.type, Afstemning.sagstrinid, Afstemning.opdateringsdato FROM Afstemning JOIN AFstemningstype ON Afstemning.typeid = Afstemningstype.id WHERE Afstemning.id=" + fmt.Sprintf("%d", *args.Id) + ";"
	
	afstemning := Afstemning{}
	var konklusion sql.NullString
	var kommentar sql.NullString
	var sagstringid sql.NullInt32
	var opdateringsdato string
	
	row := repo.db.QueryRow(query)
	
	err = row.Scan(&afstemning.Id, &afstemning.Nummer, &konklusion, &afstemning.Vedtaget, &kommentar, &afstemning.MødeID ,&afstemning.Type, &sagstringid, &opdateringsdato)

	if err != nil {
		// TODO: Implement proper graphql errors, see: https://github.com/graph-gophers/graphql-go#custom-errors
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
	
	return
}

func (a *AfstemningResolver) Id() int32 {
	return a.afstemning.Id
}