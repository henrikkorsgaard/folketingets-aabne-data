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
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestStemmeByAfstemningsIds(t *testing.T) {
	ids := []int{9351, 9352, 9353, 9354, 9355}
	repo := newRepository()
	_, err := repo.getStemmeByAfstemningIds(ids)
	assert.NoError(t, err)
}

func TestStemmeLoaderByAfstemningsIds(t *testing.T) {
	ids := []int{9351, 9352, 9353, 9354, 9355}
	for _, key := range ids {
		stemmer, err := LoadStemmerFromAfstemning(key)
		assert.NoError(t, err)
		assert.Equal(t, key, stemmer[100].AfstemningId)
	}
}
