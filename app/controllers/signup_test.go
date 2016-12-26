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

func signupSetup() {
	models.InitDB(os.Getenv("dbname"))
}

func signupTeardown() {
	models.UserDelete("foo")
	models.CloseDB()
}

func TestSignupShowOK(t *testing.T) {
	signupSetup()
	ts := httptest.NewServer(http.HandlerFunc(SignupShow))
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/signup")

	res, err := http.Get(u.String())

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	signupTeardown()
}

func TestSignupPostSuccess(t *testing.T) {
	signupSetup()

	mux := http.NewServeMux()
	mux.Handle("/signup", http.HandlerFunc(SignupPost))
	mux.Handle("/", http.HandlerFunc(Index))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/signup")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.PostForm(u.String(), url.Values{
		"username": {"foo"},
		"password": {"testpass"},
		"password_confirmation": {"testpass"},
	})

	assert.NoError(t, err)
	assert.Equal(t, 307, res.StatusCode)

	signupTeardown()
}

func TestSignupPostFail(t *testing.T) {
	signupSetup()

	mux := http.NewServeMux()
	mux.Handle("/signup", http.HandlerFunc(SignupPost))
	mux.Handle("/", http.HandlerFunc(Index))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/signup")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.PostForm(u.String(), url.Values{
		"foo": {"bar"}, // bad form post
	})

	assert.NoError(t, err)
	assert.Equal(t, "/signup", res.Request.URL.Path)
	assert.Equal(t, 200, res.StatusCode)

	signupTeardown()
}
