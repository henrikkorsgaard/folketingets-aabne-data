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

