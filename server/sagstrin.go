package server

import (
	"net/http"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

func GetSagstrinBySagsId(ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

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

			for i := range sagstrin {
				if len(sagstrin[i].Afstemning) > 0 {
					sagstrin[i].HasAfstemning = true
				}
			}

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			templateEngine.ExecuteTemplate(w, "sagstrin", sagstrin)
		},
	)
}
