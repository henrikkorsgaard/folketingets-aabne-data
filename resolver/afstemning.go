package resolver

import "fmt"

type AfstemningQueryArgs struct {
	Id *int 
}

type Afstemning struct {
	Id int32
}

func NewAfstemning(AfstemningQueryArgs) (*AfstemningResolver, error) {

	a := AfstemningResolver{} 

	repo := newSqlite()
	sql := "SELECT id FROM Afstemning LIMIT 1"
	rows, err := repo.db.Query(sql)
	defer rows.Close()
	if err != nil {
		return &a, err
	}

	var id int

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			break
		}
	}
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	return &a, nil
}

type AfstemningResolver struct {
	afstemning Afstemning
}

func (a *AfstemningResolver) Id() int32 {
	var id int32 = 10
	return id
}