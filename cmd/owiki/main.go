package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)
type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/edit/")
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
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
	r.ParseForm()
	body := r.Form.Get("body")
	(&Page{Title: title, Body: []byte(body)}).save()
	http.Redirect(w, r, "/view/" + title, 302)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := titleFromPath(r, "/view/")
	p, _ := loadPage(title)
	//log.Print(fmt.Fprintf(w, "<h1>%s</h1><div>%s</div><a href=\"/edit/" + title + "\">Edit&rsaquo;</a>",
	//		p.Title, p.Body))
	renderTemplate(w, "view", p)
}

func renderTemplate(writer http.ResponseWriter, tmpl string, page *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(writer, page)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}