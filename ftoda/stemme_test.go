package ftoda

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Println("Running tests for Stemme")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/oda.test.sqlite.db")
}

func TestStemmeLoaderByAfstemningIds(t *testing.T) {
	ids := []int{290, 291, 292, 293, 294}
	for _, key := range ids {
		stemmer, err := LoadStemmerFromAfstemning(key)
		assert.NoError(t, err)
		assert.Equal(t, key, stemmer[100].AfstemningId) //A vote always have 179 votes.
		assert.NotEmpty(t, stemmer[100].Type, "Testing that Stemme type is not empty")
	}
}

func TestStemmeLoaderByAktørIds(t *testing.T) {
	ids := []int{213,214,215,216}
	for _, key := range ids {
		stemmer, err := LoadStemmerFromAktør(key)
		assert.NoError(t, err)
		assert.Equal(t, key, stemmer[0].AktørId) //Aktør votes n times in career
		assert.NotEmpty(t, stemmer[0].Type, "Testing that Stemme type is not empty")
	}
}

