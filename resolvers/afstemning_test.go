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
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestAfstemning(t *testing.T){
	args :=AfstemningQueryArgs{}
	afstemninger, err := NewAfstemningList(args)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
}

func TestAfstemningById(t *testing.T){
	var id int32 = 8357
	args := AfstemningQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	afstemning, err := NewAfstemning(args)

	assert.NoError(t, err)
	assert.Equal(t, id, afstemning.Id())
}


func TestAfstemningNotFoundError(t *testing.T){
	var id int32 = 2
	args := AfstemningQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	_, err := NewAfstemning(args)
	assert.ErrorContains(t, err, "Unable to resolve")
}
