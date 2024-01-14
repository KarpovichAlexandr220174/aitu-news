package ui

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed html
var content embed.FS

// Templates инициализирует шаблоны
func Templates() *template.Template {
	return template.Must(template.ParseFS(content, "html/*.html"))
}

// RenderTemplate рендерит HTML-шаблон с использованием данных
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates := Templates()
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
