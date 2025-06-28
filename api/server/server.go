package server

import (
	"net/http"
	"slices"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

// Should be put in .env
var originAllowlist = []string{
	"http://127.0.0.1:8000",
	"http://localhost:8000",
}

// Pattern adopted from https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func NewServer(ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, ftodaService, templateEngine)
	var handler http.Handler = mux
	handler = checkCORS(handler)

	return handler
}

func addRoutes(mux *http.ServeMux, ftodaService *ftoda.FTODAService, templateEngine *templates.TemplateEngine) {
	mux.Handle("/lovforslag", GetLovforslag(ftodaService, templateEngine))
	mux.Handle("/lovforslag/{id}", GetLovforslagById(ftodaService, templateEngine))
	mux.Handle("/lovforslag/update", UpdateLovforslag(ftodaService, templateEngine))
}

func checkCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if slices.Contains(originAllowlist, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)

		}
		w.Header().Add("Vary", "Origin")
		next.ServeHTTP(w, r)
	})
}
