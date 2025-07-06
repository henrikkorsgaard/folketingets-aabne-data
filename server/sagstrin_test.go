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

func TestGetSagstrinBySagId(t *testing.T) {
	is := is.New(t)

	engine := templates.NewTemplateEngine()
	service := ftoda.NewFTODAService("oda.ft.dk", "../ftoda.db")
	server := NewServer(&service, &engine)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/sagstrin?sagid=101403", strings.NewReader(""))
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
	//Not the best test, but hey - it's what we got!
	is.True(strings.Contains(rows[0], "Åbent samråd med forsvarsministeren"))
}
