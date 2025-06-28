package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

func GetLovforslag(ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {
	/*
		- how do we handle pagination?
		- I can add a limit and a skip as parameters and control the page like that

	*/
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			limit, err := strconv.Atoi(r.PathValue("limit"))
			if err != nil {
				fmt.Println(err)
				limit = 0
			}
			fmt.Println(limit)
			sager, err := ftodaService.GetLovforslag(limit, 0)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			templateEngine.ExecuteTemplate(w, "list", sager)
		},
	)
}

func GetLovforslagById(ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			idString := r.PathValue("id")
			//we need to check for zero and return
			fmt.Println(idString)
			fmt.Println("does this even hit")

			id, err := strconv.Atoi(idString)
			if err != nil {
				panic(err)
			}

			//101403
			sag, err := ftodaService.GetLovforslagById(id)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			templateEngine.ExecuteTemplate(w, "lovforslag", sag)
		},
	)
}

type SagsUpdate struct {
	Count int64
	Total int64
}

func UpdateLovforslag(ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, updated, err := ftodaService.UpdateLovforslag()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500: Database update failed"))
				w.Write([]byte(err.Error()))

			}

			total := ftodaService.GetLovforslagCount()

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			templateEngine.ExecuteTemplate(w, "update", SagsUpdate{updated, total})
		},
	)
}
