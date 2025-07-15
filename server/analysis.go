package server

import (
	"net/http"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda/analysis"
	"github.com/henrikkorsgaard/folketingets-aabne-data/repository"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
)

func GetAnalysis(ftodaService *repository.FTODAService, templateEngine *templates.TemplateEngine) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			analysis.LovforslagSagstrinTypeDistribution(100, ftodaService)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte("testing analysis!"))

		},
	)
}
