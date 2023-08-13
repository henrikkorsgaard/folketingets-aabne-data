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

	/*
		The current implementation has a n+1 problem on field resolvers.
		I need to use a dataloader
		https://www.youtube.com/watch?v=uCbFMZYQbxE

		See this example: https://github.com/OscarYuen/go-graphql-starter/blob/master/server.go

		https://david-yappeter.medium.com/the-importance-of-dataloader-in-graphql-go-4d5214869b20

		https://github.com/graphql/dataloader

		Also, use context for some variables

	*/

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
