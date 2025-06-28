package templates

import (
	"html/template"
	"net/http"
	"os"
)

// See: https://kilb.tech/golang-templates - embed files solution
type TemplateEngine struct {
	tmpl *template.Template
}

func NewTemplateEngine() TemplateEngine {
	//check if the working dir is api
	//tmpl, err := template.ParseFiles("templates/lovforslag.gohtml")

	//check if the working dir is server
	tmpl, err := template.ParseFiles("../templates/lovforslag.gohtml")

	if err != nil {
		panic(err)
	}
	engine := TemplateEngine{
		tmpl,
	}
	return engine
}

func (te *TemplateEngine) ExecuteTemplate(w http.ResponseWriter, name string, data any) error {
	// we want to make sure that the templates are loaded on each request when we are developing
	if environment := os.Getenv("ENVIRONMENT"); environment == "test" {
		tmpl, err := template.ParseFiles("templates/lovforslag.gohtml")
		if err != nil {
			return err
		}

		te.tmpl = tmpl
	}

	return te.tmpl.ExecuteTemplate(w, name, data)
}
