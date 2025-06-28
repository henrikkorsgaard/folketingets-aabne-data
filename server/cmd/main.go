package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/server"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	odataHost := os.Getenv("ODATA_HOST")
	ftodaService := ftoda.NewFTODAService(odataHost, dbHost)
	templateEngine := templates.NewTemplateEngine()

	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", server.NewServer(&ftodaService, &templateEngine)))
}
