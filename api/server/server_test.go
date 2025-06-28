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
	r := httptest.NewRequest(http.MethodGet, "/lovforslag", strings.NewReader(""))
	server.ServeHTTP(w, r)

	is.Equal(w.Code, http.StatusOK)

}
