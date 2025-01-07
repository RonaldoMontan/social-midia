package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Localiza e insere os templates html na variavel 'template'
func LoadingTemplates() {

	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// Renderiza uma página html na tela
func ExecTemplate(w http.ResponseWriter, template string, data interface{}) {

	templates.ExecuteTemplate(w, template, data)	
}