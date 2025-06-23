package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/handlers"
)

func main() {
	ftodaService := ftoda.NewFTODAService()
	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", handlers.NewServer(&ftodaService)))
}
