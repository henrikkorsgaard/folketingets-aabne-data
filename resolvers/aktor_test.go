package resolvers

import (
	"os"
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	
)

func init(){
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/oda.test.sqlite.db")
}

func TestAktor(t *testing.T){
	args := AktorQueryArgs{}
	aktorer, err := NewAktorList(args)
	assert.NoError(t, err)
	assert.Len(t, aktorer, 100)
}

func TestAktorById(t *testing.T){
	var id int32 = 200
	args := AktorQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	aktor, err := NewAktor(args)

	assert.NoError(t, err)
	assert.Equal(t, id, aktor.Id())
}

func TestAktorByType(t *testing.T){
	aktortype := "Ministertitel"
	args := AktorQueryArgs{Type:&aktortype}
	aktorer, err := NewAktorList(args)
	assert.NoError(t, err)
	assert.Equal(t, aktortype, aktorer[0].Type())
}

func TestAktorByName(t *testing.T){
	name := "Karen Ellemann"
	args := AktorQueryArgs{Navn:&name}
	aktor, err := NewAktor(args)

	assert.NoError(t, err)
	assert.Equal(t, name, *aktor.Navn())
}

func TestAktorOffset(t *testing.T){
	var offset int32 = 215
	args := AktorQueryArgs{QueryArgs:QueryArgs{Offset:&offset}}
	aktorer, err := NewAktorList(args)
	assert.NoError(t, err)
	assert.Equal(t, aktorer[0].Id(), int32(216))
}

func TestAktorLimit(t *testing.T){
	var limit int32 = 215
	args := AktorQueryArgs{QueryArgs:QueryArgs{Limit:&limit}}
	aktorer, err := NewAktorList(args)

	assert.NoError(t, err)
	assert.Len(t, aktorer, 215)
}


func TestAktorNotFoundError(t *testing.T){
	var id int32 = 20000
	args := AktorQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	_, err := NewAktor(args)
	assert.ErrorContains(t, err, "unable to resolve")
}

func TestSearchAktorByName(t *testing.T) {
	name := "An"
	args := AktorSearchArgs{Navn:name}
	_, err := NewAktorResultList(args)
	assert.NoError(t, err)
	// Hard to do any additional asserts here.
}

func TestAktorRelation(t *testing.T) {
	name := "Eva Esmarch"
	args := AktorQueryArgs{Navn:&name}
	aktor, err := NewAktor(args)
	rels, err := aktor.Relationer()
	
	for _, r := range *rels {
		fmt.Printf("%+v\n", *r.Relation())
	}
	fmt.Printf("%+v\n", rels)
	assert.NoError(t, err)
	assert.Equal(t, name, *aktor.Navn())
}
