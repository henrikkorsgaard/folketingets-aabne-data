package resolvers

import (
	"context"
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

func TestStemme(t *testing.T) {

	args := StemmeQueryArgs{QueryArgs: QueryArgs{}}
	_, err := NewStemmeList(context.Background(),args)
	assert.NoError(t, err)
}

func TestStemmeById(t *testing.T) {
	var id int32 = 2129580
	args := StemmeQueryArgs{QueryArgs: QueryArgs{Id: &id}}
	_, err := NewStemmeList(context.Background(),args)

	assert.NoError(t, err)
}

func TestStemmeByAfstemningId(t *testing.T) {
	var afsid int32 = 9351
	args := StemmeQueryArgs{AfstemningId: &afsid}
	_, err := NewStemmeList(context.Background(),args)

	assert.NoError(t, err)
}

func TestStemmeByAfstemningId2(t *testing.T) {
	var afsid int32 = 9351
	args := StemmeQueryArgs{AfstemningId: &afsid}
	_, err := NewStemme(context.Background(),args)

	assert.NoError(t, err)
}

func TestStemmeByIdAndAfstemningId(t *testing.T) {
	var id int32 = 2129580
	var afsid int32 = 9351
	args := StemmeQueryArgs{AfstemningId: &afsid, QueryArgs: QueryArgs{Id: &id}}
	_, err := NewStemmeList(context.Background(),args)

	assert.NoError(t, err)
}

func TestStemmeNotFoundError(t *testing.T) {
	var id int32 = 2
	args := StemmeQueryArgs{QueryArgs: QueryArgs{Id: &id}}
	_, err := NewStemmeList(context.Background(),args)
	assert.ErrorContains(t, err, "record not found")
}
