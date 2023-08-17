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

func TestAktørLoaderAll(t *testing.T) {
	aktører, err := LoadAktører(100, 0)
	assert.NoError(t, err)
	assert.Len(t, aktører, 100)
	assert.NotEmpty(t, aktører[0].Type, "Testing that Aktør.Type is not empty")
}

func TestAktørTypeJoin(t *testing.T) {
	aktør, err := LoadAktørById(1528)
	assert.NoError(t, err)
	assert.Equal(t,"Privatperson", aktør.Type)
}

func TestAktørLoaderByAktørType(t *testing.T) {
	aktørtype := "Ministertitel"
	aktører, err := LoadAktørerByType(100, 0, aktørtype)
	assert.NoError(t, err)
	assert.Equal(t, aktørtype, aktører[0].Type)
}

func TestAktørLoadByName(t *testing.T) {
	name := "Karen Ellemann"
	aktør, err := LoadAktørByName(name)
	assert.NoError(t, err)
	assert.Equal(t, name, aktør.Navn)
}

func TestAktørLoadByIds(t *testing.T) {
	ids := []int{107,136,140,142,150,165}
	for _, key := range ids {
		aktør, err := LoadAktørById(key)
		assert.NoError(t, err)
		assert.Equal(t, key, aktør.Id)
	}
}

// We wan to make sure the loader will return errors on ids not found
func TestAktørNotFoundError(t *testing.T){
	idError := 200000
	_, err := LoadAktørById(idError)
	assert.ErrorContains(t, err, "record not found")
}

// We wan to make sure the loader will return errors on ids not found when we send multiple request to the loader
func TestAktørFoundThenNotFoundError(t *testing.T){
	idExist := 107
	idError := 200000

	aktør, err := LoadAktørById(idExist)
	assert.NoError(t, err)
	assert.Equal(t, idExist, aktør.Id)
	
	_, err = LoadAktørById(idError)
	assert.ErrorContains(t, err, "record not found")
}

func TestSearchAktørByName(t *testing.T) {
	name := "An"
	_, err := SearchAktørByName(100,name)
	assert.NoError(t, err)
	// Hard to do any additional asserts here.
}

func TestSearchAktørByNamePrivateIndividual(t *testing.T) {
	name := "Anja Lund"
	aktører, err := SearchAktørByName(100,name)
	
	assert.NoError(t, err)
	assert.Len(t, aktører, 1)
}
