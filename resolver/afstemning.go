package resolver

import "fmt"

type AfstemningQueryArgs struct {
	Id *int 
}

type Afstemning struct {
	Id int32
}

func NewAfstemning(AfstemningQueryArgs) (*AfstemningResolver, error) {
	newSqlite()
	fmt.Println("is this even called?")
	a := AfstemningResolver{} 
	return &a, nil
}

type AfstemningResolver struct {
	afstemning Afstemning
}

func (a *AfstemningResolver) Id() int32 {
	var id int32 = 10
	return id
}