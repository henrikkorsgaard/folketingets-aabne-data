package resolver

import (
	"fmt"
	"testing"
	

	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	

)

func init(){
	fmt.Println("Running tests for the Afstemning")
	godotenv.Load("../config_dev.env")
}

func TestAfstemning(t *testing.T){

	assert.Equal(t, true, false)
}