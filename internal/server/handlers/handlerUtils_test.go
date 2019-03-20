package handlers_test

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/oletizi/owiki/internal/server/handlers"
	"log"
	"testing"
)

var (
	bodyText = "I'm the body"
	j        = []byte(fmt.Sprintf(`{"body": "%s"}`, bodyText))
)

func TestDecodeJsonGeneric(t *testing.T) {
	var f interface{}
	err := handlers.DecodeJson(j, &f)
	if err != nil {
		t.Error(err)
	}

	m := f.(map[string]interface{})

	assert.Equal(t, bodyText, m["body"])
}

func TestDecodeJsonStruct(t *testing.T) {
	type JsonBody struct {
		Body string
	}

	var b = JsonBody{}
	err := handlers.DecodeJson(j, &b)
	if err != nil {
		t.Error(err)
	}

	log.Print("body text: " + b.Body)

	assert.Equal(t, bodyText, b.Body)
}
