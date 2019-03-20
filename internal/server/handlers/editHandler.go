package handlers

import (
	"github.com/oletizi/owiki/internal/page"
	"log"
	"net/http"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := TitleFromPath(r, "/edit/")
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
		log.Print(err)
	}
	RenderTemplate(w, "edit", p)
}
