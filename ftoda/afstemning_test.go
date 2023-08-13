package ftoda

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	
)

func init(){
	fmt.Println("Running tests for Afstemning")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestAfstemningById(t *testing.T){
	var id int32 = 8357
	repo := NewRepository()
	afstemning, err := repo.GetAfstemning(id)
	assert.NoError(t, err)
	assert.Equal(t, id, afstemning.Id)
}