package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/henrikkorsgaard/folketingets-aabne-data/ftoda"
	"github.com/henrikkorsgaard/folketingets-aabne-data/server"
	"github.com/henrikkorsgaard/folketingets-aabne-data/templates"
	"github.com/matryer/is"
)

func TestGetLovforslagLimit(t *testing.T) {
	is := is.New(t)

	engine := templates.NewTemplateEngine()
	service := ftoda.NewFTODAService("oda.ft.dk", "../ftoda.db")
	server := server.NewServer(&service, &engine)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/lovforslag?limit=1", strings.NewReader(""))
	server.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK) //expect statuscode 200

	temp := strings.Split(w.Body.String(), "\n")
	rows := []string{}
	for _, s := range temp {
		s = strings.TrimSpace(s)
		if s != "" {
			rows = append(rows, s)
		}
	}

	is.Equal(len(rows), 1) //expect length to be 1
}

func TestGetLovforslag(t *testing.T) {
	is := is.New(t)

	engine := templates.NewTemplateEngine()
	service := ftoda.NewFTODAService("oda.ft.dk", "../ftoda.db")
	server := server.NewServer(&service, &engine)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/lovforslag", strings.NewReader(""))
	server.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK) //expect statuscode 200

	temp := strings.Split(w.Body.String(), "\n")
	rows := []string{}
	for _, s := range temp {
		s = strings.TrimSpace(s)
		if s != "" {
			rows = append(rows, s)
		}
	}

	is.Equal(len(rows), 100) //expect length to be 100
}

func TestGetLovforslagById(t *testing.T) {
	is := is.New(t)

	engine := templates.NewTemplateEngine()
	service := ftoda.NewFTODAService("oda.ft.dk", "../ftoda.db")
	server := server.NewServer(&service, &engine)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/lovforslag/101403", strings.NewReader(""))
	server.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK) //expect statuscode 200
	is.True(strings.Contains(w.Body.String(), "<h1>Forslag til lov om forsvarssamarbejde mellem Danmark og Amerikas Forenede Stater m.v.</h1>"))
}
