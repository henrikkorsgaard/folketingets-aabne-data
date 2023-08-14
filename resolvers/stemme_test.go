package resolvers

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Println("Running tests for the Afstemning")
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/odatest.sqlite.db")
}

func TestStemmerByAfstemningId(t *testing.T) {
	var afsid int32 = 9351
	args := StemmeQueryArgs{AfstemningId: &afsid}
	stemmeResolvers, err := NewStemmeList(args)

	assert.NoError(t, err)
	assert.Equal(t, int(afsid), stemmeResolvers[0].stemme.AfstemningId)
}
