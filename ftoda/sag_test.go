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

func TestSagLoadAll(t *testing.T) {
	sag, err := LoadSager(100,0)
	assert.NoError(t, err)
	assert.Len(t, sag, 100)
	assert.NotEmpty(t, sag[0].Type, "Testing that Sag.Type is not empty")
	assert.NotEmpty(t, sag[0].Kategori, "Testing that Sag.Kategori is not empty")
	assert.NotEmpty(t, sag[0].Status, "Testing that Sag.Status is not empty")
}

func TestSagLoadByIds(t *testing.T){
	ids := []int{100,101,102,103,104,105}
	for _, id := range ids {
		sag, err := LoadSag(id)
		assert.NoError(t, err)
		assert.Equal(t, id, sag.Id)
	}
}
