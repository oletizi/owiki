package handlers

import (
	"github.com/oletizi/owiki/internal/page"
	"io/ioutil"
	"net/http"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	type BodyText struct {
		Body string
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, err)
	}

	var bodyText BodyText
	DecodeJson(body, &bodyText)
	title := TitleFromPath(r, "/save/")
	p := &page.Page{Title: title, Body: []byte(bodyText.Body)}
	err = p.Save()

	if err != nil {
		HandleError(w, err)
		return
	}
	http.Redirect(w, r, "/view/"+title, 302)
}
