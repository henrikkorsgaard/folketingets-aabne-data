package resolvers

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

func TestStemmerByAfstemningId(t *testing.T) {
	var afsid int32 = 206
	args := StemmeQueryArgs{AfstemningId: &afsid}
	stemmeResolvers, err := NewStemmeList(args)

	assert.NoError(t, err)
	assert.Equal(t, int(afsid), stemmeResolvers[0].stemme.AfstemningId)
}
