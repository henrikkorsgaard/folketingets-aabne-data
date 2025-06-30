package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"

	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
	"github.com/matryer/is"
)

func TestAfstemningFromSagId(t *testing.T) {
	is := is.New(t)

	engine := templates.NewTemplateEngine()
	service := ftoda.NewFTODAService("oda.ft.dk", "../ftoda.db")
	server := NewServer(&service, &engine)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/afstemning?sagstrinid=266904", strings.NewReader(""))
	server.ServeHTTP(w, r)

	is.Equal(w.Code, http.StatusOK) //expect statuscode 200
	is.True(strings.Contains(w.Body.String(), "<h1>10353</h1>"))
}
