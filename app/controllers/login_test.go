package controllers

import (
	"testing"
	"net/http"
	"net/url"
	"os"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"bytes"

	"github.com/JacksonGariety/cetch/app/models"
)

func loginSetup() {
	models.InitDB(os.Getenv("dbname"))
}

func loginTeardown() {
	models.UserDelete("foo")
	models.CloseDB()
}

func TestLoginShowOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(LoginShow))
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/login")

	res, err := http.Get(u.String())

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

func TestLoginPostSuccess(t *testing.T) {
	loginSetup()
	models.UserCreate("foo", "testpass")

	mux := http.NewServeMux()
	mux.Handle("/login", http.HandlerFunc(LoginPost))
	mux.Handle("/", http.HandlerFunc(Index))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/login")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.PostForm(u.String(), url.Values{
		"username": {"foo"},
		"password": {"testpass"},
	})

	assert.NoError(t, err)
	assert.Equal(t, 307, res.StatusCode)

	loginTeardown()
}

func TestLoginPostFail(t *testing.T) {
	loginSetup()

	mux := http.NewServeMux()
	mux.Handle("/login", http.HandlerFunc(LoginPost))
	mux.Handle("/", http.HandlerFunc(Index))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/login")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.PostForm(u.String(), url.Values{
		"foo": {"bar"}, // bad form post
	})

	assert.NoError(t, err)
	assert.Equal(t, "/login", res.Request.URL.Path)
	assert.Equal(t, 200, res.StatusCode)

	loginTeardown()
}
