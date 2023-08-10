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
	
	args := AfstemningQueryArgs{}
	_, err := NewAfstemningList(args)
	assert.NoError(t, err)
}

func TestAfstemningById(t *testing.T){
	var id int32 = 8357
	args := AfstemningQueryArgs{&id}
	_, err := NewAfstemningList(args)

	assert.NoError(t, err)
}


func TestAfstemningNotFoundError(t *testing.T){
	var id int32 = 2
	args := AfstemningQueryArgs{&id}
	_, err := NewAfstemning(args)
	assert.ErrorContains(t, err, "Unable to resolve Afstemning")
}
