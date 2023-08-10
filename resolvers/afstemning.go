package resolvers

import (
	"database/sql"
	"fmt"
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
	//MødeID int32 // Table: Møde
	Type string // Table: Afstemningstype 
	//SagstringID integer // Table: Sagstrin
	//Opdateringsdato Time.time // DB: as text
}

func NewAfstemning(args AfstemningQueryArgs) (*AfstemningResolver, error) {
	
	resolver := AfstemningResolver{} 
	repo := newSqlite()

	// SELECT Afstemning.*, Afstemningstype.type FROM Afstemning JOIN Afstemningstype ON Afstemning.typeid = Afstemningstype.id;

	query := "SELECT Afstemning.id, Afstemning.nummer, Afstemning.konklusion, Afstemning.vedtaget, Afstemning.kommentar, Afstemningstype.type FROM Afstemning JOIN AFstemningstype ON Afstemning.typeid = Afstemningstype.id WHERE Afstemning.id=" + fmt.Sprintf("%d", 8474) + ";"
	
	afstemning := Afstemning{}
	var konklusion sql.NullString
	var kommentar sql.NullString
	
	row := repo.db.QueryRow(query)
	
	err := row.Scan(&afstemning.Id, &afstemning.Nummer, &konklusion, &afstemning.Vedtaget, &kommentar,&afstemning.Type)
	if err != nil {
		fmt.Println(err.Error())
	}

	if konklusion.Valid {
		afstemning.Konklusion = konklusion.String
	}

	if kommentar.Valid {
		afstemning.Kommentar = kommentar.String
	}
	
	return &resolver, nil
}

func (a *AfstemningResolver) Id() int32 {
	return a.afstemning.Id
}