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

func TestAfstemningAll(t *testing.T){
	repo := NewRepository()
	afstemninger, err := repo.GetAllAfstemning(200, 0)
	assert.NoError(t, err)
	assert.Len(t, afstemninger,200)
}

func TestAfstemningByIdList(t *testing.T){
	repo := NewRepository()
	ids := []int{8370,8371,8372,8373,8374,8375}
	afstemninger, err := repo.GetAfstemningByIds(ids)
	assert.NoError(t, err)
	assert.Len(t, afstemninger,len(ids))
}

func TestAfstemningById(t *testing.T){
	var id int = 8357
	repo := NewRepository()
	afstemning, err := repo.GetAfstemning(id)
	assert.NoError(t, err)
	assert.Equal(t, id, afstemning.Id)
}