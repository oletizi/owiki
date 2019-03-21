package handler_test

import (
	"github.com/bxcodec/faker"
	"github.com/magiconair/properties/assert"
	"github.com/oletizi/owiki/internal/page"
	"github.com/oletizi/owiki/internal/server/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	rr         *httptest.ResponseRecorder
	handleEdit http.HandlerFunc
	tmpDir     string
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setup() {
	tmp, err := ioutil.TempDir("/tmp/", "test")
	if err != nil {
		panic(err)
	}

	editTemplate := tmp + "/edit.gohtml"
	includesTemplate := tmp + "/includes.gohtml"

	// TODO: handle error
	ioutil.WriteFile(editTemplate, []byte("I'm the edit template."), 0644)
	ioutil.WriteFile(includesTemplate, []byte("I'm the includes template."), 0644)

	rr = httptest.NewRecorder()
	config := handler.Config{
		Docroot:      tmp,
		TemplateRoot: tmp,
		PageFactory:  page.NewPage,
	}
	h := handler.NewHandler(&config)
	handleEdit = http.HandlerFunc(h.HandleEdit)
}

func teardown() {
	os.RemoveAll(tmpDir)
}

func TestHandleEditNoTitle(t *testing.T) {
	// Test no title specified
	req, err := http.NewRequest("GET", "/view/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handleEdit.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusNotFound)
}

func TestHandleEditWithTitle(t *testing.T) {
	title := faker.Word()
	req, err := http.NewRequest("GET", "/view/"+title, nil)
	if err != nil {
		t.Fatal(err)
	}

	handleEdit.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}
