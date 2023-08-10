package resolvers

import (
	"fmt"
	"strings"
)

type AktørResolver struct {
	aktør Aktør
}

type Aktør struct {
	Id int32
}

func NewAktørList(args QueryArgs) (resolvers []*AktørResolver, err error){

	if args.Id != nil {
		aktørResolver, err := NewAktør(args)
		if aktørResolver != nil {
			resolvers = append(resolvers, aktørResolver)
		}
		return resolvers, err
	}

	repo := newSqlite()

	query := "SELECT Aktør.id FROM Aktør;"

	rows, err := repo.db.Query(query)
	if err != nil {
		return
	}

	for rows.Next(){
		aktør := Aktør{}

		err = rows.Scan(&aktør.Id)

		if err != nil {
			break
		}

		aktørResolver := AktørResolver{aktør}
		resolvers = append(resolvers, &aktørResolver)
	}

	if rows.Err() != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			err = fmt.Errorf("Unable to resolve Afstemning in database.")
		}
	}

	return
}

func NewAktør(args QueryArgs) (resolver *AktørResolver,err error) {
	repo := newSqlite()

	query := "SELECT Aktør.id FROM Aktør WHERE Aktør.id=" + fmt.Sprintf("%d", *args.Id) + ";"
	
	aktør := Aktør{}
	
	row := repo.db.QueryRow(query)
	
	err = row.Scan(&aktør.Id)

	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			err = fmt.Errorf("Unable to resolve Aktør: Id %d does not exist", *args.Id)
		}
		return
	}

	resolver = &AktørResolver{aktør}

	return 
}

func (a *AktørResolver) Id() int32 {
	return a.aktør.Id
}
