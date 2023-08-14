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

func TestAfstemningLoaderAll(t *testing.T) {
	afstemninger, err := LoadAfstemninger(100, 0)
	assert.NoError(t, err)
	assert.Len(t, afstemninger, 100)
}

func TestAfstemningLoadByIds(t *testing.T) {
	ids := []int{9351, 9352, 9353, 9354, 9355}
	for _, key := range ids {
		afstemning, err := LoadAfstemning(key)
		assert.NoError(t, err)
		assert.Equal(t, key, afstemning.Id)
	}
}
