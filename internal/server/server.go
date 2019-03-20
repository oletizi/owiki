package server

import (
	"log"
	"net/http"
	"strconv"
)

func Run(port int) {

	portString := strconv.Itoa(port)

	log.Print("Configuring server...")
	http.HandleFunc("/view/", ViewHandler)
	http.HandleFunc("/edit/", EditHandler)
	http.HandleFunc("/save/", SaveHandler)
	log.Print("Starting server on port " + portString)
	log.Fatal(http.ListenAndServe(":"+portString, nil))
}
