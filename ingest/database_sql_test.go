package utils

import (
	"fmt"
	"testing"
	"os"
	
	"io/ioutil"
	"database/sql"
	
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/microsoft/go-mssqldb"
)

func init(){
	fmt.Println("Running database setup and connection tests.")
	godotenv.Load("../config_dev.env")
}

func TestPSQLCreateDatabase(t *testing.T) {

	sqlFilePath := "./sql/oda.psql.sql"

	user := os.Getenv("PSQL_DATABASE_USER")
	pass := os.Getenv("PSQL_DATABASE_PASS")
	name := os.Getenv("PSQL_DATABASE_NAME")
	host := os.Getenv("PSQL_DATABASE_HOST")
	port := os.Getenv("PSQL_DATABASE_PORT")

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

func TestSqliteCreateDatabase(t *testing.T) {

	sqlFilePath := "./sql/oda.sqlite.sql"

	db, err := sql.Open("sqlite3", "./sql/oda.sqlite.db")
	
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

func TestSqliteTestData(t *testing.T) {

	db, err := sql.Open("sqlite3", "./sql/odatest.sqlite.db")
	defer db.Close()
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	row, err := db.Query("SELECT id FROM Afstemning LIMIT 1")
	assert.NoError(t, err)
	defer row.Close()

	// When we start developing for real, we need to unmarshal into struct
	var id int

	for row.Next() {
		row.Scan(&id)
	}

	assert.NotEmpty(t, id)
}

func TestMSSQLDatabaseAccess(t *testing.T) {

	user := os.Getenv("MSSQL_DATABASE_USER")
	pass := os.Getenv("MSSQL_DATABASE_PASS")
	name := os.Getenv("MSSQL_DATABASE_NAME")
	host := os.Getenv("MSSQL_DATABASE_HOST")
	port := os.Getenv("MSSQL_DATABASE_PORT")

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s/%s", user,pass,host,port,name)

	db, err := sql.Open("sqlserver", connectionString)
	
	defer db.Close()
	if err != nil {
		t.Error("Unexpected error in test " + t.Name() + ": " + err.Error())
	}

	_, err = db.Exec("SELECT @@VERSION")
	assert.NoError(t, err)
}