package server

import (
	"github.com/oletizi/owiki/page"
	"html/template"
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

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/save/")
	body := r.Form.Get("body")
	err := (&page.Page{Title: title, Body: []byte(body)}).Save()
	if err != nil {
		handleError(w, err)
	}
	http.Redirect(w, r, "/view/"+title, 302)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/view/")
	p, _ := page.LoadPage(title)
	renderTemplate(w, "view", p)
}

func renderTemplate(writer http.ResponseWriter, tmpl string, page *page.Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		handleError(writer, err)
		return
	}
	err = t.Execute(writer, page)
	if err != nil {
		handleError(writer, err)
	}
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
