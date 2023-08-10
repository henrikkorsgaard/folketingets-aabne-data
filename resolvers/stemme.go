package resolvers
import (
	
	"fmt"
	"time"

	graphql "github.com/graph-gophers/graphql-go"

)

type StemmeResolver struct {
	stemme Stemme
}

type Stemme struct {
	Id int32
	Type string
	AfstemningID int32
	AktørID int32 
	Opdateringsdato graphql.Time
}

func NewStemmeList(args QueryArgs) (resolvers []*StemmeResolver, err error){

	repo := newSqlite()

	query := "SELECT Stemme.id, Stemmetype.type, Stemme.afstemningid, Stemme.aktørid, Stemme.opdateringsdato FROM Stemme JOIN Stemmetype ON Stemme.typeid = Stemmetype.id;"

	if args.Id != nil {
		query = query[0:len(query)-1]
		query +=  " WHERE Stemme.id=" + fmt.Sprintf("%d", *args.Id) + ";"
	}

	rows, err := repo.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next(){
		stemme := Stemme{}
		var opdateringsdato string

		err = rows.Scan(&stemme.Id, &stemme.Type, &stemme.AfstemningID, &stemme.AktørID, &opdateringsdato)

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

func NewStemme(args QueryArgs) (resolver *StemmeResolver,err error) {
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
	args := QueryArgs{&s.stemme.AktørID}
	return NewAktør(args)
}

func (s *StemmeResolver) Afstemning() (*AfstemningResolver, error) {
	args := QueryArgs{&s.stemme.AfstemningID}
	return NewAfstemning(args)
}
