package main

import (
        "log"
        "net/http"
		"io/ioutil"
		"os"

        graphql "github.com/graph-gophers/graphql-go"
        "github.com/graph-gophers/graphql-go/relay"
		"github.com/joho/godotenv"

		"henrikkorsgaard/folketingets-aabne-data/resolvers"
)

func main() {

	godotenv.Load("../config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "ingest/data/odatest.sqlite.db")

	b, err := ioutil.ReadFile("./schema/schema.graphql")
	if err != nil {
		panic(err)
	}

	qr := resolvers.QueryResolver{}

    schema := graphql.MustParseSchema(string(b), &qr)
	http.HandleFunc("/", handleTestClient)
    http.Handle("/query", &relay.Handler{Schema: schema})
	
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTestClient(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"client/client.html")
}