package resolvers

import (
	
	"fmt"
	"strings"
	
)

type MødeQueryArgs struct {
	Id *int32
}

type MødeResolver struct {
	møde Møde
}

type Møde struct {
	Id int32
}

func NewMøde(args MødeQueryArgs) (resolver *MødeResolver,err error) {
	
	repo := newSqlite()

	query := "SELECT Møde.id FROM Møde WHERE Møde.id=" + fmt.Sprintf("%d", *args.Id) + ";"
	
	møde := Møde{}
	
	row := repo.db.QueryRow(query)
	
	err = row.Scan(&møde.Id)

	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			err = fmt.Errorf("Unable to resolve Møde: Id %d does not exist", *args.Id)
		}
		return
	}

	resolver = &MødeResolver{møde}

	return 
}

func (mr *MødeResolver) Id() int32 {
	return mr.møde.Id
}