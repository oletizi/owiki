package handler_test

import (
	"github.com/bxcodec/faker"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/oletizi/owiki/internal/mock/mock_page"
	"github.com/oletizi/owiki/internal/page"
	"github.com/oletizi/owiki/internal/server/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	rr         *httptest.ResponseRecorder
	handleEdit http.HandlerFunc
	handleSave http.HandlerFunc
	tmpDir     string
	config     handler.Config
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

	// TODO:
	// - handle errors
	// - consider whether to mock the template interaction; maybe overkill
	ioutil.WriteFile(editTemplate, []byte("I'm the edit template."), 0644)
	ioutil.WriteFile(includesTemplate, []byte("I'm the includes template."), 0644)

	rr = httptest.NewRecorder()
	config = handler.Config{
		Docroot:      tmp,
		TemplateRoot: tmp,
		PageFactory:  page.NewPage,
	}
	h := handler.NewHandler(&config)
	handleEdit = http.HandlerFunc(h.HandleEdit)
	handleSave = http.HandlerFunc(h.HandleSave)
}

func teardown() {
	os.RemoveAll(tmpDir)
}

// Test no title specified
func TestHandler_HandleViewNoTitle(t *testing.T) {
	ctrl, p := setupMockPage(t)
	defer ctrl.Finish()

	// return
	p.EXPECT().LoadPage().Return(nil)

	req, err := http.NewRequest("GET", "/view/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handleEdit.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusNotFound)
}

// Test title specified
func TestHandler_HandleViewWithTitle(t *testing.T) {
	ctrl, p := setupMockPage(t)
	defer ctrl.Finish()

	p.EXPECT().LoadPage()

	title := faker.Word()
	req, err := http.NewRequest("GET", "/view/"+title, nil)
	if err != nil {
		t.Fatal(err)
	}

	handleEdit.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestHandler_HandleSaveNoTitle(t *testing.T) {
	// test no title
	ctrl, p := setupMockPage(t)
	defer ctrl.Finish()

	// expect save not to be called when there's no title
	p.EXPECT().Save().Times(0)

	req, err := http.NewRequest("POST", "/save/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handleSave.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusNotFound)
}

func TestHandler_HandleSaveNoBody(t *testing.T) {
	ctrl, p := setupMockPage(t)
	defer ctrl.Finish()

	p.EXPECT().Save()

	title := faker.Word()
	req, err := http.NewRequest("POST", "/save/"+title, nil)
	if err != nil {
		t.Fatal(err)
	}

	handleSave.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusFound)
	assert.Equal(t, p.Title(), title)
	assert.Equal(t, p.Body(), []byte(""))
}

func TestHandler_HandleSaveWithTitleAndBody(t *testing.T) {
	ctrl, p := setupMockPage(t)
	defer ctrl.Finish()

	p.EXPECT().Save()

	title := faker.Word()
	body := "the body"
	req, err := http.NewRequest("POST", "/save/"+title, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	handleSave.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusFound)
	assert.Equal(t, p.Title(), title)
	assert.Equal(t, p.Body(), []byte(body))
}

func setupMockPage(t *testing.T) (*gomock.Controller, *mock_page.MockPage) {
	ctrl := gomock.NewController(t)
	p := mock_page.NewMockPage(ctrl)
	config.PageFactory = func(title string, body []byte) page.Page {

		p.EXPECT().Title().AnyTimes().Return(title)
		p.EXPECT().Body().AnyTimes().Return(body)
		return p
	}
	return ctrl, p
}
