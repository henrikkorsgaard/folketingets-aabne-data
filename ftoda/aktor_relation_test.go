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

func TestLoadRelationsFromAktor(t *testing.T) {
	rels, err := LoadAktorRelations(914)
	assert.NoError(t, err)
	for _, rel := range rels {
		assert.Equal(t, 914, rel.FraAktorId)
	}
	
}

func TestLoadRelationsFromAktorer(t *testing.T){
	ids := []int{1,6,20,146,914}
	for _, id := range ids {
		rels, err := LoadAktorRelations(id)
		assert.NoError(t, err)
		for _, rel := range rels {
			assert.Equal(t, id, rel.FraAktorId)
		}
	}
}
