package templates

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:embed *.gohtml
var filesystem embed.FS

type TemplateEngine struct {
	tmpl *template.Template
}

func NewTemplateEngine() TemplateEngine {

	tmpl, err := template.ParseFS(filesystem, "lovforslag.gohtml", "afstemning.gohtml")
	if err != nil {
		panic(err)
	}
	engine := TemplateEngine{
		tmpl,
	}
	return engine
}

/*
Proxy function that allow us to load templates dynamically
on dev environment.
*/
func (te *TemplateEngine) ExecuteTemplate(w http.ResponseWriter, name string, data any) error {
	// we want to make sure that the templates are loaded on each request when we are developing
	if environment := os.Getenv("ENVIRONMENT"); environment == "dev" {
		fmt.Println("Dev environment: Parsing temlate on every load")
		tmpl, err := template.ParseFiles("templates/lovforslag.gohtml", "templates/afstemning.gohtml")
		if err != nil {
			return err
		}

		te.tmpl = tmpl
	}

	return te.tmpl.ExecuteTemplate(w, name, data)
}
