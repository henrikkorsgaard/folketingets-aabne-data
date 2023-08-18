package resolvers


import (
	"os"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	
)

func init(){
	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "../ingest/data/oda.test.sqlite.db")
}

func TestSagerAll(t *testing.T){
	args := SagQueryArgs{}
	sager, err := NewSagList(args)
	assert.NoError(t, err)
	assert.Len(t, sager, 100)
}

func TestSagerByType(t *testing.T){
	sagType := "Lovforslag"
	args := SagQueryArgs{Type:&sagType}
	sager, err := NewSagList(args)
	assert.NoError(t, err)
	for _, s := range sager {
		assert.Equal(t, sagType, *s.Type())
	}
}

func TestSagById(t *testing.T) {
	var id int32 = 10
	args := SagQueryArgs{QueryArgs:QueryArgs{Id:&id}}
	sag, err := NewSagList(args)
	assert.NoError(t, err)
	assert.Equal(t, id, sag[0].Id())
}
