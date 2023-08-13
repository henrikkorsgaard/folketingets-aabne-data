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
	
	args := StemmeQueryArgs{QueryArgs: QueryArgs{}}
	_, err := NewStemmeList(args)
	assert.NoError(t, err)
}

func TestStemmeById(t *testing.T){
	var id int32 = 2129580
	args := StemmeQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	_, err := NewStemmeList(args)

	assert.NoError(t, err)
}

func TestStemmeByAfstemningId(t *testing.T){
	var afsid int32 = 9351
	args := StemmeQueryArgs{AfstemningId:&afsid}
	_, err := NewStemmeList(args)

	assert.NoError(t, err)
}

func TestStemmeByIdAndAfstemningId(t *testing.T){
	var id int32 = 2129580
	var afsid int32 = 9351
	args := StemmeQueryArgs{AfstemningId: &afsid,QueryArgs:QueryArgs{Id:&id}}
	_, err := NewStemmeList(args)

	assert.NoError(t, err)
}

func TestStemmeNotFoundError(t *testing.T){
	var id int32 = 2
	args := StemmeQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	_, err := NewStemmeList(args)
	assert.ErrorContains(t, err, "Unable to resolve Stemme")
}
