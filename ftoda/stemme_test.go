package ftoda

import (
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
	afstemninger, err := repo.GetAllStemme(200, 0)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 200)
}

func TestStemmeByIdList(t *testing.T) {
	repo := NewRepository()
	ids := []int{2129581, 2129582, 2129583, 2129584, 2129585}
	afstemninger, err := repo.GetStemmeByIds(ids)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, len(ids))
}

func TestStemmeById(t *testing.T) {
	var id int = 2129580
	repo := NewRepository()
	afstemning, err := repo.GetStemme(id)
	assert.NoError(t, err)
	assert.Equal(t, id, afstemning.Id)
}
