package handlers

import (
	"github.com/oletizi/owiki/internal/page"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := TitleFromPath(r, "/view/")
	p, err := page.LoadPage(title)
	if err != nil {
		HandleError(w, err)
	} else {
		RenderTemplate(w, "view", p)
	}
}
