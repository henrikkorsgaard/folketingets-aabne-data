package ftoda

import (
	"fmt"
	"os"

	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Println("Running tests for Aktør")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestAktørLoaderAll(t *testing.T) {
	aktører, err := LoadAfstemninger(100, 0)
	assert.NoError(t, err)
	assert.Len(t, aktører, 100)
}

func TestAktørTypeJoin(t *testing.T) {
	aktør, err := LoadAktørById(19109)
	assert.NoError(t, err)
	assert.Equal(t,"Privatperson", aktør.Type)
}

func TestAktørLoadByName(t *testing.T) {
	name := "Anne Madsen"
	aktør, err := LoadAktørByName(name)
	assert.NoError(t, err)
	assert.Equal(t, name, aktør.Navn)
}

func TestAktørLoadByIds(t *testing.T) {
	ids := []int{19107,19108,19109,19110,19111}
	for _, key := range ids {
		aktør, err := LoadAktørById(key)
		assert.NoError(t, err)
		assert.Equal(t, key, aktør.Id)
	}
}

// We wan to make sure the loader will return errors on ids not found
func TestAktørNotFoundError(t *testing.T){
	idError := 2
	_, err := LoadAktørById(idError)
	assert.ErrorContains(t, err, "record not found")
}

// We wan to make sure the loader will return errors on ids not found when we send multiple request to the loader
func TestAktørFoundThenNotFoundError(t *testing.T){
	idExist := 19109
	idError := 2

	aktør, err := LoadAktørById(idExist)
	assert.NoError(t, err)
	assert.Equal(t, idExist, aktør.Id)
	
	_, err = LoadAktørById(idError)
	assert.ErrorContains(t, err, "record not found")
}
