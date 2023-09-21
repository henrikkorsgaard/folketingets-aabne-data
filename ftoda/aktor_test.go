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

func TestAktorLoaderAll(t *testing.T) {
	aktorer, err := LoadAktorer(100, 0)
	assert.NoError(t, err)
	assert.Len(t, aktorer, 100)
	assert.NotEmpty(t, aktorer[0].Type, "Testing that Aktor.Type is not empty")
}

func TestAktorTypeJoin(t *testing.T) {
	aktor, err := LoadAktorById(1528)
	assert.NoError(t, err)
	assert.Equal(t,"Privatperson", aktor.Type)
}

func TestAktorLoaderByAktorType(t *testing.T) {
	aktortype := "Ministertitel"
	aktorer, err := LoadAktorerByType(100, 0, aktortype)
	assert.NoError(t, err)
	assert.Equal(t, aktortype, aktorer[0].Type)
}

func TestAktorLoadByName(t *testing.T) {
	name := "Karen Ellemann"
	aktor, err := LoadAktorByName(name)
	assert.NoError(t, err)
	assert.Equal(t, name, aktor.Navn)
}

func TestAktorLoadByIds(t *testing.T) {
	ids := []int{107,136,140,142,150,165}
	for _, key := range ids {
		aktor, err := LoadAktorById(key)
		assert.NoError(t, err)
		assert.Equal(t, key, aktor.Id)
	}
}

// We wan to make sure the loader will return errors on ids not found
func TestAktorNotFoundError(t *testing.T){
	idError := 200000
	_, err := LoadAktorById(idError)
	assert.ErrorContains(t, err, "record not found")
}

// We wan to make sure the loader will return errors on ids not found when we send multiple request to the loader
func TestAktorFoundThenNotFoundError(t *testing.T){
	idExist := 107
	idError := 200000

	aktor, err := LoadAktorById(idExist)
	assert.NoError(t, err)
	assert.Equal(t, idExist, aktor.Id)
	
	_, err = LoadAktorById(idError)
	assert.ErrorContains(t, err, "record not found")
}

func TestSearchAktorByName(t *testing.T) {
	name := "An"
	_, err := SearchAktorByName(100,name)
	assert.NoError(t, err)
	// Hard to do any additional asserts here.
}

func TestSearchAktorByNamePrivateIndividual(t *testing.T) {
	name := "Anja Lund"
	aktorer, err := SearchAktorByName(100,name)
	
	assert.NoError(t, err)
	assert.Len(t, aktorer, 1)
}
