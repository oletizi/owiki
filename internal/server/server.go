package server

import (
	"github.com/oletizi/owiki/internal/page"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/edit/")
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
		log.Print(err)
	}
	renderTemplate(w, "edit", p)
}

func titleFromPath(r *http.Request, prefix string) string {
	title := r.URL.Path[len(prefix):]
	return title
}

type body struct {
	body string
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("in save handler...")
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
	}

	log.Print("body:")
	log.Print(body)

	title := titleFromPath(r, "/save/")
	log.Print("title from path: " + title)

	p := &page.Page{Title: title}

	err = p.Save()

	if err != nil {
		handleError(w, err)
		return
	}
	http.Redirect(w, r, "/view/"+title, 302)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/view/")
	p, err := page.LoadPage(title)
	if err != nil {
		handleError(w, err)
	} else {
		renderTemplate(w, "view", p)
	}
}

func renderTemplate(writer http.ResponseWriter, tmpl string, page *page.Page) {
	t, err := template.ParseFiles(getTemplatePath(tmpl), getTemplatePath("includes"))
	if err != nil {
		handleError(writer, err)
		return
	}
	err = t.Execute(writer, page)
	if err != nil {
		handleError(writer, err)
	}
}

func getTemplatePath(tmpl string) string {
	return "web/" + tmpl + ".gohtml"
}

func handleError(writer http.ResponseWriter, err error) {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}

func Run(port int) {

	portString := strconv.Itoa(port)

	log.Print("Configuring server...")
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Print("Starting server on port " + portString)
	log.Fatal(http.ListenAndServe(":"+portString, nil))
}
