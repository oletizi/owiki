package page

import (
	"io/ioutil"
	"log"
)

type Factory interface {
	NewPage(title string, body []byte) Page
}

type Page interface {
	Save() error
	LoadPage() error
	Title() string
	Body() []byte
}

type filePageFactory struct {
	config *FilePageConfig
}

type FilePageConfig struct {
	DataDir string
}

func NewFilePageFactory(config *FilePageConfig) Factory {
	f := filePageFactory{
		config: config,
	}
	return &f
}

func (f filePageFactory) NewPage(title string, body []byte) Page {
	p := filePage{
		title:  title,
		body:   body,
		config: f.config,
	}
	return &p
}

type filePage struct {
	config *FilePageConfig
	title  string
	body   []byte
}

func (p *filePage) Save() error {
	filename := "/tmp/" + p.Title() + ".txt"
	log.Print("saving file: " + filename)
	log.Print("body text: " + string(p.body))
	return ioutil.WriteFile(filename, p.body, 0600)
}

func (p *filePage) LoadPage() error {
	filename := "/tmp/" + p.title + ".txt"
	log.Print("loading file: " + filename)
	body, err := ioutil.ReadFile(filename)
	log.Print(err)
	p.body = body
	return err
}

func (p *filePage) Title() string {
	return p.title
}

func (p *filePage) Body() []byte {
	return p.body
}
