package server

import (
	"net/http"
	"strconv"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda/lovforslag"
	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

func GetLovforslag(ftodaService *repository.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {
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

			sager, err := ftodaService.GetSagerByType(3, limit, offset)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			err = templateEngine.ExecuteTemplate(w, "list", sager)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		},
	)
}

/*
We should include sagstin for lovforslag. This includes the history of the legislation
//https://oda.ft.dk/api/Sagstrin?$format=json&$filter=sagid%20eq%20102266
*/
func GetLovforslagById(ftodaService *repository.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.Atoi(r.PathValue("id"))

			if err != nil {
				panic(err)
			}

			//101403
			lovforslag, err := lovforslag.NewFromSagId(id, ftodaService)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				w.Write([]byte(err.Error()))
			}

			/*
				for _, emneordsag := range sag.EmneordSager {
					emne, err := ftodaService.GetEmneordById(emneordsag.EmneordId)
					if err == nil {
						sag.Emneord = append(sag.Emneord, emne.Emneord)
					}
				}
				fmt.Println(sag.Emneord)
			*/
			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			err = templateEngine.ExecuteTemplate(w, "lovforslag", lovforslag)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		},
	)
}

// There is a lot of abstraction potential here as soon as I add emneord etc.
type SagsUpdate struct {
	Count int64
	Total int64
}

/*
func UpdateLovforslag(ftodaService *repository.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, updated, err := ftodaService.UpdateSagerByType(3)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500: Database update failed"))
				w.Write([]byte(err.Error()))

			}

			total := ftodaService.GetLovforslagCount()

			//TODO: Set headers globally with a proxy handler
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			err = templateEngine.ExecuteTemplate(w, "update", SagsUpdate{updated, total})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		},
	)
}*/
