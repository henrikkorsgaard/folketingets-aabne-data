package resolvers
import (
	
	"fmt"
	"time"
	"os"

	graphql "github.com/graph-gophers/graphql-go"

)

type StemmeQueryArgs struct {
	QueryArgs
	AfstemningId *int32
}

type StemmeResolver struct {
	stemme Stemme
}

type Stemme struct {
	Id int32
	Type string
	AfstemningId int32
	AktørID int32 
	Opdateringsdato graphql.Time
}

func NewStemmeList(args StemmeQueryArgs) (resolvers []*StemmeResolver, err error){

	repo := newSqlite()

	query := "SELECT Stemme.id, Stemmetype.type, Stemme.afstemningid, Stemme.aktørid, Stemme.opdateringsdato FROM Stemme JOIN Stemmetype ON Stemme.typeid = Stemmetype.id"

	if args.Id != nil {
		query +=  " WHERE Stemme.id=" + fmt.Sprintf("%d", *args.Id)
	}

	if args.AfstemningId != nil {
		
		afstemningId := fmt.Sprintf("%d", *args.AfstemningId)
		if args.Id != nil {
			query += " AND Stemme.afstemningid=" + afstemningId
		} else {
			query +=  " WHERE Stemme.afstemningid=" + afstemningId 
		}
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
		stemme := Stemme{}
		var opdateringsdato string

		err = rows.Scan(&stemme.Id, &stemme.Type, &stemme.AfstemningId, &stemme.AktørID, &opdateringsdato)

		if err != nil {
			break
		}

		t, err := time.Parse(time.DateTime, opdateringsdato)
		if err != nil {
			break
		}

		stemme.Opdateringsdato.Time = t

		stemmeResolver := StemmeResolver{stemme}
		resolvers = append(resolvers, &stemmeResolver)
	}

	if rows.Err() != nil {
		err = rows.Err()
	}

	if args.Id != nil && len(resolvers) == 0 {
		err = fmt.Errorf("Unable to resolve Stemme: Id %d does not exist", *args.Id)
	}

	return
}

func NewStemme(args StemmeQueryArgs) (resolver *StemmeResolver,err error) {
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
	return s.stemme.Id
}

func (s *StemmeResolver) Type() *string {
	return &s.stemme.Type
}

func (s *StemmeResolver) Opdateringsdato() graphql.Time {
	return s.stemme.Opdateringsdato
}

func (s *StemmeResolver) Aktør() (*AktørResolver, error) {
	args := AktørQueryArgs{QueryArgs:QueryArgs{Id: &s.stemme.AktørID}}
	return NewAktør(args)
}

func (s *StemmeResolver) Afstemning() (*AfstemningResolver, error) {
	args := QueryArgs{Id:  &s.stemme.AfstemningId}
	return NewAfstemning(args)
}
