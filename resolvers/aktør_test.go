package resolvers

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	
)

func init(){
	fmt.Println("Running tests for the Aktør")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestAktør(t *testing.T){
	args := AktørQueryArgs{}
	aktører, err := NewAktørList(args)
	assert.NoError(t, err)
	assert.Len(t, aktører, 100)
}

func TestAktørById(t *testing.T){
	var id int32 = 20029
	args := AktørQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	aktør, err := NewAktør(args)

	assert.NoError(t, err)
	assert.Equal(t, id, aktør.Id())
}

func TestAktørByType(t *testing.T){
	aktørtype := "Ministertitel"
	args := AktørQueryArgs{Type:&aktørtype}
	aktører, err := NewAktørList(args)
	assert.NoError(t, err)
	assert.Equal(t, aktørtype, aktører[0].Type())
}

func TestAktørByName(t *testing.T){
	name := "Anne Madsen"
	args := AktørQueryArgs{Navn:&name}
	aktør, err := NewAktør(args)

	assert.NoError(t, err)
	assert.Equal(t, name, *aktør.Navn())
}


func TestAktørNotFoundError(t *testing.T){
	var id int32 = 2
	args := AktørQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	_, err := NewAktør(args)
	assert.ErrorContains(t, err, "unable to resolve")
}
