package ftoda

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Println("Running tests for Afstemning")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestStemmeAll(t *testing.T) {
	repo := NewRepository()
	stemmer, err := repo.GetAllStemme(200, 0)
	assert.NoError(t, err)
	assert.Len(t, stemmer, 200)
}

func TestStemmeByIdList(t *testing.T) {
	repo := NewRepository()
	ids := []int{2129581, 2129582, 2129583, 2129584, 2129585}
	stemmer, err := repo.GetStemmeByIds(ids)
	assert.NoError(t, err)
	assert.Len(t, stemmer, len(ids))
}

func TestStemmeById(t *testing.T) {
	var id int = 2129580
	repo := NewRepository()
	stemme, err := repo.GetStemme(id)
	assert.NoError(t, err)
	assert.Equal(t, id, stemme.Id)
}

func TestStemmeByAfstemningsIds(t *testing.T) {
	ids := []int{9351, 9352, 9353, 9354, 9355}
	repo := NewRepository()
	_, err := repo.GetStemmeByAfstemningIds(ids)
	assert.NoError(t, err)
	
}

func TestStemmeLoaderByAfstemningsIds(t *testing.T) {
	ids := []int{9351, 9352, 9353, 9354, 9355}
	loader := NewStemmeLoader()
	for _, key := range ids {
		thunk := loader.Load(context.Background(), key)
		result, err := thunk()
		assert.NoError(t, err)
		stemmer := *result
		assert.Equal(t, key, stemmer[100].AfstemningId)
	}
}
