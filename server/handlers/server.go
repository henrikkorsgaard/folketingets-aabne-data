package handlers

import (
	"net/http"
	"slices"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
)

// Should be put in .env
var originAllowlist = []string{
	"http://127.0.0.1:8000",
	"http://localhost:8000",
}

// Pattern adopted from https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
func NewServer(ftodaService *ftoda.FTODAService) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, ftodaService)
	var handler http.Handler = mux
	handler = checkCORS(handler)

	return handler
}

func addRoutes(mux *http.ServeMux, ftodaService *ftoda.FTODAService) {
	mux.Handle("/lovforslag", GetLovforslag(ftodaService))
	mux.Handle("/lovforslag/{id}", GetLovforslag(ftodaService))
	mux.Handle("/lovforslag/update", UpdateLovforslag(ftodaService))
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
