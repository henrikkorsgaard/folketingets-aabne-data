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

/*
See: https://github.com/pacedotdev/oto/blob/main/otohttp/server_test.go
See: https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
*/
func TestServer(t *testing.T) {
	is := is.New(t)

	engine := templates.NewTemplateEngine()
	service := ftoda.NewFTODAService("oda.ft.dk", "../ftoda.db")
	server := NewServer(&service, &engine)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	server.ServeHTTP(w, r)

	is.Equal(w.Code, http.StatusNotFound) // expect statuscode 404
}

func TestGetLovforslagLimit(t *testing.T) {
	is := is.New(t)

	engine := templates.NewTemplateEngine()
	service := ftoda.NewFTODAService("oda.ft.dk", "../ftoda.db")
	server := NewServer(&service, &engine)

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
	server := NewServer(&service, &engine)

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
