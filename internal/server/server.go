package server

import (
	"github.com/oletizi/owiki/internal/page"
	"github.com/oletizi/owiki/internal/server/handler"
	"log"
	"net/http"
	"strconv"
)

func Run(port int) {

	portString := strconv.Itoa(port)

	log.Print("Configuring server...")

	config := handler.Config{
		Docroot:      "./web",
		TemplateRoot: "./web",
		PageFactory:  page.NewPage,
	}

	h := handler.NewHandler(&config)

	http.HandleFunc("/view/", h.HandleView)
	http.HandleFunc("/edit/", h.HandleEdit)
	http.HandleFunc("/save/", h.HandleSave)
	log.Print("Starting server on port " + portString)
	log.Fatal(http.ListenAndServe(":"+portString, nil))
}
