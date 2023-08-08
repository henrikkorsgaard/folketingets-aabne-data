package repository


import (
	"fmt"
	"testing"
	"os"
	
	"io/ioutil"
	"database/sql"

	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init(){
	fmt.Println("Running database setup tests")
	godotenv.Load("../config_dev.env")
}

func TestPSQLCreateDatabase(t *testing.T) {

	sqlFilePath := "./sql/oda.psql.sql"

	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user,pass,host,port,name)
	db, err := sql.Open("postgres", connectionString)
	
	defer db.Close()
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	file, err := ioutil.ReadFile(sqlFilePath)
    if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
    }

	_, err = db.Exec(string(file))
	assert.NoError(t, err)
}
/*
func TestSqliteCreateDatabase(t *testing.T) {

	sqlFilePath := "./sql/oda.sqlite.sql"

	pool, err := pgxpool.New(context.Background(), connectionString)
	defer pool.Close()
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	file, err := ioutil.ReadFile(sqlFilePath)
    if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
    }

	_, err = pool.Exec(context.Background(), string(file))
	assert.NoError(t, err)
}*/