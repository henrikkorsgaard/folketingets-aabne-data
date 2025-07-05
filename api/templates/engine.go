package templates

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:embed **/*.gohtml
var folder embed.FS

type TemplateEngine struct {
	tmpl *template.Template
}

func NewTemplateEngine() TemplateEngine {

	//need to add the file structure from in here.
	//components
	//pages
	tmpl, err := template.ParseFS(folder, "*/*.gohtml")
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
		tmpl, err := template.ParseGlob("templates/*/*.gohtml")
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("%+v", tmpl)
		te.tmpl = tmpl
	}

	return te.tmpl.ExecuteTemplate(w, name, data)
}
