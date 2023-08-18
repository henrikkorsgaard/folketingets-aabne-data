package resolvers

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	
)

func init(){
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/oda.test.sqlite.db")
}

func TestAfstemningAll(t *testing.T){
	args :=AfstemningQueryArgs{}
	afstemninger, err := NewAfstemningList(args)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
}

func TestAfstamningAllWithKommentar(t *testing.T){
	hasComments := true
	args := AfstemningQueryArgs{Kommentar:&hasComments}
	afstemninger, err := NewAfstemningList(args)
	assert.NoError(t, err)
	for _, a := range afstemninger {
		assert.NotEmpty(t, *a.Kommentar())
	}
}

func TestAfstemningByType(t *testing.T){
	afstemningsType := "Endelig vedtagelse"
	args :=AfstemningQueryArgs{Type:&afstemningsType}
	afstemninger, err := NewAfstemningList(args)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
}

func TestAfstemningById(t *testing.T){
	var id int32 = 1
	args := AfstemningQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	afstemning, err := NewAfstemning(args)

	assert.NoError(t, err)
	assert.Equal(t, id, afstemning.Id())
}

func TestAfstemningNotFoundError(t *testing.T){
	var id int32 = 20000
	args := AfstemningQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	_, err := NewAfstemning(args)
	assert.ErrorContains(t, err, "Unable to resolve")
}

