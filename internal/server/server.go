package server

import (
	"github.com/oletizi/owiki/internal/page"
	"github.com/oletizi/owiki/internal/server/handler"
	"log"
	"net/http"
	"strconv"
)

func Run(port int, dataDir string, docroot string, templateDir string) {

	log.Print("Configuring server...")
	pageConfig := &page.FilePageConfig{
		DataDir: dataDir,
	}
	pageFactory := page.NewFilePageFactory(pageConfig)

	portString := strconv.Itoa(port)

	handlerConfig := handler.Config{
		Docroot:      docroot,
		TemplateRoot: templateDir,
		PageFactory:  pageFactory,
	}

	h := handler.NewHandler(&handlerConfig)

	http.HandleFunc("/view/", h.HandleView)
	http.HandleFunc("/edit/", h.HandleEdit)
	http.HandleFunc("/save/", h.HandleSave)
	log.Print("Starting server on port " + portString)
	log.Fatal(http.ListenAndServe(":"+portString, nil))
}
