package server

import (
	"net/http"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

func GetLovforslag(ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			q := r.URL.Query()

			limit, err := strconv.Atoi(q.Get("limit"))
			if err != nil {
				limit = 0
			}

			offset, err := strconv.Atoi(q.Get("offset"))
			if err != nil {
				offset = 0
			}

			sager, err := ftodaService.GetLovforslag(limit, offset)
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
			q := r.URL.Query()
			id, err := strconv.Atoi(q.Get("id"))
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
