package server

import (
	"net/http"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

func GetAfstemningBySagstrinId(ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			sagstrinid, err := strconv.Atoi(q.Get("sagstrinid"))
			if err != nil {
				panic(err)
			}

			afstemning, err := ftodaService.GetAfstemningBySagstrinId(sagstrinid)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}
			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			templateEngine.ExecuteTemplate(w, "afstemning", afstemning)
		},
	)
}
