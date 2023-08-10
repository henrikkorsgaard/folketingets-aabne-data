package resolvers

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	
)

func init(){
	fmt.Println("Running tests for the Stemme")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestStemme(t *testing.T){
	
	args := StemmeQueryArgs{}
	_, err := NewStemmeList(args)
	assert.NoError(t, err)
}

func TestStemmeById(t *testing.T){
	var id int32 = 8357
	args := StemmeQueryArgs{&id}
	_, err := NewStemeList(args)

	assert.NoError(t, err)
}


func TestStemmeNotFoundError(t *testing.T){
	var id int32 = 2
	args := StemmeQueryArgs{&id}
	_, err := NewStemmeList(args)
	assert.ErrorContains(t, err, "Unable to resolve Afstemning")
}
