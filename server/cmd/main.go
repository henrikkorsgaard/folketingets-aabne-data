package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/handlers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbFile := os.Getenv("DB_FILE")
	odataHost := os.Getenv("ODATA_HOST")
	ftodaService := ftoda.NewFTODAService(odataHost, dbFile)
	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", handlers.NewServer(&ftodaService)))
}
