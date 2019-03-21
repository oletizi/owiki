package handler

import "github.com/oletizi/owiki/internal/page"

type Config struct {
	Docroot      string
	TemplateRoot string
	PageFactory  func(title string, body []byte) page.Page
}
