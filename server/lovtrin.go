package server

import (
	"net/http"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

func GetLovtrinBySagsId(ftodaService *repository.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			sagid, err := strconv.Atoi(q.Get("sagid"))
			if err != nil {
				panic(err)
			}

			sagstrin, err := ftodaService.GetSagstrinBySagsId(sagid)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			templateEngine.ExecuteTemplate(w, "sagstrin", sagstrin)
		},
	)
}

/*
func GetLovforslagsTrinBySagsId(ftodaService *repository.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			sagid, err := strconv.Atoi(q.Get("sagid"))
			if err != nil {
				panic(err)
			}

			sagstrin, err := ftodaService.GetSagstrinBySagsId(sagid)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}

			//sagstrin is array
			lovforslagstrin := ftoda.InitiateLovforslagsTrin(sagid)
			for _, st := range sagstrin {

				switch st.Typeid {
				case 6, 20, 19, 31, 77:
				default:
					lovforslagstrin[7].Sagstrin = append(lovforslagstrin[7].Sagstrin, st)
				}
			}

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			templateEngine.ExecuteTemplate(w, "sagstrin", sagstrin)
		},
	)
}*/

func GetlovtrinById(ftodaService *repository.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.Atoi(r.PathValue("id"))
			if err != nil {
				panic(err)
			}

			sagstrin, err := ftodaService.GetSagstrinById(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			templateEngine.ExecuteTemplate(w, "sagstrin-details", sagstrin)
		},
	)
}
