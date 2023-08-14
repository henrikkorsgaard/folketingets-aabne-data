package resolvers

import (
	
	
)

type MødeResolver struct {
	møde Møde
}

type Møde struct {
	Id int32
}

func NewMøde(args QueryArgs) (resolver *MødeResolver,err error) {
	/*
	repo := newSqlite()

	query := "SELECT Møde.id FROM Møde WHERE Møde.id=" + fmt.Sprintf("%d", *args.Id)

	if args.Offset == nil {
		var offset int32 = 0
		args.Offset = &offset
	}

	query+= " LIMIT " + os.Getenv("GRAPHQL_QUERY_LIMIT") + " OFFSET " + fmt.Sprintf("%d", *args.Offset) + ";"
	
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
	*/
	return 
}

func (mr *MødeResolver) Id() int32 {
	return mr.møde.Id
}