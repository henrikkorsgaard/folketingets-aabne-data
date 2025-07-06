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

	//"2025-09-04T00:00:00"
	//ts := "2025-09-04T00:00:00"
	/*
		//If I need to use the date time, then I need to add my own marshall step

		see: https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json

		ts := "2006-01-02T15:04:05Z07:00"
		theTime, err := time.Parse("2006-01-02T03:04:05Z", ts)
		if err != nil {
			fmt.Println("Could not parse time:", err)
		}
		fmt.Println(theTime)
	*/

	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", server.NewServer(&ftodaService, &templateEngine)))
}
