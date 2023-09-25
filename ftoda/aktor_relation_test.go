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

func TestAktorLoadAktorRelations(t *testing.T) {
	_, err := LoadAktorRelations(914)
	assert.NoError(t, err)
}
