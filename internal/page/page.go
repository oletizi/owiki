package page

import (
	"io/ioutil"
	"log"
)

type Page interface {
	Save() error
	LoadPage() error
	Title() string
	Body() []byte
}

func NewPage(title string, body []byte) Page {
	p := page{
		title: title,
		body:  body,
	}
	return &p
}

type page struct {
	title string
	body  []byte
}

func (p *page) Save() error {
	filename := "/tmp/" + p.Title() + ".txt"
	log.Print("saving file: " + filename)
	log.Print("body text: " + string(p.body))
	return ioutil.WriteFile(filename, p.body, 0600)
}

func (p *page) LoadPage() error {
	filename := "/tmp/" + p.title + ".txt"
	log.Print("loading file: " + filename)
	body, err := ioutil.ReadFile(filename)
	log.Print(err)
	p.body = body
	return err
}

func (p *page) Title() string {
	return p.title
}

func (p *page) Body() []byte {
	return p.body
}
