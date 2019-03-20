package page

import (
	"io/ioutil"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := "/tmp/" + p.Title + ".txt"
	log.Print("saving file: " + filename)
	log.Print("body text: " + string(p.Body))
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := "/tmp/" + title + ".txt"
	log.Print("loading file: " + filename)
	body, err := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}, err
}
