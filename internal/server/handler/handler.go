package handler

import (
	"encoding/json"
	"github.com/oletizi/owiki/internal/page"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Handler struct {
	config *Config
}

func NewHandler(config *Config) *Handler {
	handler := Handler{config: config}
	return &handler
}

func (h *Handler) HandleEdit(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/edit/")
	if title == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	p := h.config.PageFactory(title, nil)

	// XXX: ignores the error; should probably have a mechanism for deciding if the page *should* be there and handle
	// the error explicitly instead of implicitly assuming that an error loading the page means the doesn't exist yet.
	p.LoadPage()

	h.renderTemplate(w, "edit", &p)
}

func (h *Handler) HandleSave(w http.ResponseWriter, r *http.Request) {
	type BodyText struct {
		Body string
	}
	var body = []byte(`{ "body": "" }`)
	var err error
	var bodyText BodyText

	title := titleFromPath(r, "/save/")
	if title == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Body != nil {
		defer r.Body.Close()
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			handleError(w, err)
		}
	}

	decodeJson(body, &bodyText)

	log.Print("decoded body: " + bodyText.Body)

	p := h.config.PageFactory(title, []byte(bodyText.Body))
	err = p.Save()

	if err != nil {
		handleError(w, err)
		return
	}
	http.Redirect(w, r, "/view/"+title, 302)
}

func (h *Handler) HandleView(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/view/")
	p := h.config.PageFactory(title, nil)
	h.renderTemplate(w, "view", &p)
}

func (h *Handler) getTemplatePath(tmpl string) string {
	return h.config.TemplateRoot + "/" + tmpl + ".gohtml"
}

func handleError(writer http.ResponseWriter, err error) {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}

func decodeJson(data []byte, decoded interface{}) error {
	return json.Unmarshal(data, decoded)
}

func titleFromPath(r *http.Request, prefix string) string {
	title := r.URL.Path[len(prefix):]
	return title
}

func (h *Handler) renderTemplate(writer http.ResponseWriter, tmpl string, page *page.Page) {
	t, err := template.ParseFiles(h.getTemplatePath(tmpl), h.getTemplatePath("includes"))
	if err != nil {
		log.Print("current working directory: ")
		log.Print(os.Getwd())
		log.Print("Failed to load template files.")
		log.Print(err)
		handleError(writer, err)
		return
	}
	err = t.Execute(writer, page)
	if err != nil {
		handleError(writer, err)
	}
}
