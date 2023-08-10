package resolvers

import (
	"fmt"
	"testing"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	graphql "github.com/graph-gophers/graphql-go"

)

func init(){
	fmt.Println("Running tests for the query resolver")
	godotenv.Load("../config_dev.env")
}

func TestResolverSatisfySchema(t *testing.T){
	// This does not call anything, it just test if the implementation satisfy the interfaces implied by the schema
	b, err := ioutil.ReadFile("../schema/schema.graphql")
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	qr := QueryResolver{}
	_,err = graphql.ParseSchema(string(b), &qr)

	assert.NoError(t, err)
}