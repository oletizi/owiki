package handlers

import (
	"encoding/json"
	"github.com/oletizi/owiki/internal/page"
	"html/template"
	"net/http"
)

func GetTemplatePath(tmpl string) string {
	return "web/" + tmpl + ".gohtml"
}

func HandleError(writer http.ResponseWriter, err error) {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}

func DecodeJson(data []byte, decoded interface{}) error {
	return json.Unmarshal(data, decoded)
}

func TitleFromPath(r *http.Request, prefix string) string {
	title := r.URL.Path[len(prefix):]
	return title
}

func RenderTemplate(writer http.ResponseWriter, tmpl string, page *page.Page) {
	t, err := template.ParseFiles(GetTemplatePath(tmpl), GetTemplatePath("includes"))
	if err != nil {
		HandleError(writer, err)
		return
	}
	err = t.Execute(writer, page)
	if err != nil {
		HandleError(writer, err)
	}
}
