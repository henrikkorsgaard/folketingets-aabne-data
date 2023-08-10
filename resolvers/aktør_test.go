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
	args := QueryArgs{}
	_, err := NewAktørList(args)
	assert.NoError(t, err)
}

func TestAktørById(t *testing.T){
	var id int32 = 19050
	args := QueryArgs{&id}
	_, err := NewAktørList(args)

	assert.NoError(t, err)
}

func TestAktørNotFoundError(t *testing.T){
	var id int32 = 10
	args := QueryArgs{&id}
	_, err := NewAktørList(args)
	assert.ErrorContains(t, err, "Unable to resolve Aktør")
}
