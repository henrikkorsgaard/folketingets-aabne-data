package ftoda

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/oda.test.sqlite.db")
}

func TestAfstemningAll(t *testing.T) {
	afstemninger, err := LoadAfstemninger(100, 0)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
	assert.NotEmpty(t, afstemninger[0].Type, "Testing Afstemning.Type not empty")
	assert.NotEmpty(t, afstemninger[0].Vedtaget, "Testing Afstemning.Vedtaget not empty")
	assert.NotEmpty(t, afstemninger[0].ModeId, "Testing Afstemning.Mode not empty")
}

func TestAfstemningHasKommentar(t *testing.T) {
	afstemninger, err := LoadAfstemningerWithKommentar(100, 0)
	assert.NoError(t, err)
	for _, a := range afstemninger {
		assert.NotEmpty(t, a.Kommentar)
	}
}

func TestAfstemningByType(t *testing.T) {
	afstemninger, err := LoadAfstemningerByType(100, 0, "Endelig vedtagelse")
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
}

func TestAfstemningTypeJoin(t *testing.T) {
	afstemning, err := LoadAfstemning(9351)
	assert.NoError(t, err)
	assert.Equal(t,"Endelig vedtagelse", afstemning.Type)
}

func TestAfstemningLoadByIds(t *testing.T) {
	ids := []int{9351, 9352, 9353, 9354, 9355}
	for _, key := range ids {
		afstemning, err := LoadAfstemning(key)
		assert.NoError(t, err)
		assert.Equal(t, key, afstemning.Id)
	}
}

func TestAfstemningerBySagId(t *testing.T) {
	id := 1143
	afstemninger, err := LoadAfstemningerBySag(id)
	assert.NoError(t, err)
	assert.Equal(t, id, afstemninger[0].SagId)
	/*
	for _, afs := range afstemninger {
		assert.Equal(t, id, afs.SagId)
	}*/
}

// We wan to make sure the loader will return errors on ids not found
func TestAfstemningNotFoundError(t *testing.T){
	idError := 20000
	_, err := LoadAfstemning(idError)
	assert.ErrorContains(t, err, "record not found")
}

// We wan to make sure the loader will return errors on ids not found when we send multiple request to the loader
func TestAfstemningFoundThenNotFoundError(t *testing.T){
	idExist := 8357
	idError := 20000
	
	afstemning, err := LoadAfstemning(idExist)
	assert.NoError(t, err)
	assert.Equal(t, idExist, afstemning.Id)
	
	_, err = LoadAfstemning(idError)
	assert.ErrorContains(t, err, "record not found")
}

