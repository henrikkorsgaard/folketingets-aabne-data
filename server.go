package main

import (
	
	"io/ioutil"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"

	"henrikkorsgaard/folketingets-aabne-data/resolvers"

	"github.com/friendsofgo/graphiql"
)

func main() {


	godotenv.Load("config_dev.env")
	os.Setenv("SQLITE_DATABASE_PATH", "ingest/data/odatest.sqlite.db")

	b, err := ioutil.ReadFile("./schema/schema.graphql")
	if err != nil {
		panic(err)
	}

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	qr := resolvers.QueryResolver{}

	schema := graphql.MustParseSchema(string(b), &qr, graphql.MaxDepth(5))
	http.Handle("/", graphiqlHandler)
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
