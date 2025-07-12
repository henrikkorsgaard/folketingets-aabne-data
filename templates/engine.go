package templates

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

//go:embed **/*.gohtml
var folder embed.FS

type TemplateEngine struct {
	tmpl *template.Template
}

var functions = template.FuncMap{
	"join": join,
}

func join(arr []string) string {
	return strings.Join(arr[:], ", ")
}

func NewTemplateEngine() TemplateEngine {

	//tmpl, err := template.ParseFS(folder, "*/*.gohtml")
	tmpl, err := template.New("").Funcs(functions).ParseFS(folder, "*/*.gohtml")
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
		tmpl, err := template.New("").Funcs(functions).ParseGlob("templates/*/*.gohtml")
		tmpl.Funcs(functions)
		if err != nil {
			fmt.Println(err)
			return err
		}
		te.tmpl = tmpl
	}

	return te.tmpl.ExecuteTemplate(w, name, data)
}
