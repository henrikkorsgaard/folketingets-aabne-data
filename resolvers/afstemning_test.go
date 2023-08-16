package resolvers

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	
)

func init(){
	fmt.Println("Running tests for the Afstemning")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/oda.test.sqlite.db")
}

func TestAfstemningAll(t *testing.T){
	args :=AfstemningQueryArgs{}
	afstemninger, err := NewAfstemningList(args)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
	// Want to make sure not null types are in the data
	assert.NotEmpty(t, afstemninger[0].Type(), "Testing Afstemning.Type not empty")
	assert.NotEmpty(t, afstemninger[0].Vedtaget(), "Testing Afstemning.Vedtaget not empty")
	assert.NotEmpty(t, afstemninger[0].Møde(), "Testing Afstemning.Vedtaget not empty")
}

func TestAfstemningByType(t *testing.T){
	afstemningsType := "Endelig vedtagelse"
	args :=AfstemningQueryArgs{Type:&afstemningsType}
	afstemninger, err := NewAfstemningList(args)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
	// Want to make sure not null types are in the data
}

func TestAfstemningById(t *testing.T){
	var id int32 = 1
	args := AfstemningQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	afstemning, err := NewAfstemning(args)

	assert.NoError(t, err)
	assert.Equal(t, id, afstemning.Id())
	// Want to make sure not null types are in the data
	assert.NotEmpty(t, afstemning.Type(), "Testing Afstemning.Type not empty")
	assert.NotEmpty(t, afstemning.Vedtaget(), "Testing Afstemning.Vedtaget not empty")
	assert.NotEmpty(t, afstemning.Møde(), "Testing Afstemning.Møde not empty")
}


func TestAfstemningNotFoundError(t *testing.T){
	var id int32 = 20000
	args := AfstemningQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	_, err := NewAfstemning(args)
	assert.ErrorContains(t, err, "Unable to resolve")
}

