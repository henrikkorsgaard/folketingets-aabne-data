package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
)

var templateDirPath = ""

func init() {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	templateDirPath = wd + "/templates/"
	fmt.Println("Serving templates from " + templateDirPath)
}

func GetLovforslag(ftodaService *ftoda.FTODAService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// use thing to handle request
			tmpl, err := template.ParseFiles(templateDirPath + "lovforslag.gohtml")
			if err != nil {
				panic(err)
			}

			sager, err := ftodaService.GetLovforslag(0)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			tmpl.ExecuteTemplate(w, "list", sager)
		},
	)
}

func GetLovforslagById(ftodaService *ftoda.FTODAService) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			idString := r.PathValue("id")
			//we need to check for zero and return
			fmt.Println(idString)

			id, err := strconv.Atoi(idString)
			if err != nil {
				panic(err)
			}

			tmpl, err := template.ParseFiles(templateDirPath + "lovforslag.gohtml")
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

			tmpl.ExecuteTemplate(w, "lovforslag", sag)
		},
	)
}

type SagsUpdate struct {
	Count int64
	Total int64
}

func UpdateLovforslag(ftodaService *ftoda.FTODAService) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, updated, err := ftodaService.UpdateLovforslag()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500: Database update failed"))
				w.Write([]byte(err.Error()))

			}

			total := ftodaService.GetLovforslagCount()

			tmpl, err := template.ParseFiles(templateDirPath + "/lovforslag.gohtml")
			if err != nil {
				panic(err)
			}

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			tmpl.ExecuteTemplate(w, "update", SagsUpdate{updated, total})
		},
	)
}
