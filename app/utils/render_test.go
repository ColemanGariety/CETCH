package utils

import (
	"testing"
	"net/http/httptest"
	"encoding/xml"
)

type body struct {
	Content string `xml:",innerxml"`
}

type html struct {
	Body body `xml:"body"`
}

func TestRenderNoProps(t *testing.T) {
	w := httptest.NewRecorder()
	Render(w, "index.html", nil)

	h := html{}
	err := xml.NewDecoder(w.Body).Decode(&h)
	if err != nil {
		t.Error(err)
	}
}

func TestRenderProps(t *testing.T) {
	w := httptest.NewRecorder()
	Render(w, "index.html", &Props{
		"foo": "bar",
	})

	h := html{}
	err := xml.NewDecoder(w.Body).Decode(&h)
	if err != nil {
		t.Error(err)
	}
}
