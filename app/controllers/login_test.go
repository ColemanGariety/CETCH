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

func loginTestSetup() {
	models.InitDB(os.Getenv("dbname"))
}

func loginTestTeardown() {
	(&models.Users{}).DeleteAll()
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

func TestLoginNonexistentUsername(t *testing.T) {
	loginTestSetup()

	data := url.Values{ "username": {"foo"}, "password": {"bar"} }
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/login",  bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	LoginPost(w, r)

	assert.Contains(t, w.Body.String(), "<input type=\"text\" name=\"username\" value=\"foo\" />")
	assert.Contains(t, w.Body.String(), "<input type=\"password\" name=\"password\" value=\"bar\" />")
	assert.Contains(t, w.Body.String(), "invalid username or password")
	loginTestTeardown()
}

func TestLoginIncorrectPassword(t *testing.T) {
	loginTestSetup()
	(&models.User{ Name: "foo" }).CreateFromPassword("notbar")

	data := url.Values{ "username": {"foo"}, "password": {"bar"} }
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/login",  bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	LoginPost(w, r)

	assert.Contains(t, w.Body.String(), "<input type=\"text\" name=\"username\" value=\"foo\" />")
	assert.Contains(t, w.Body.String(), "<input type=\"password\" name=\"password\" value=\"bar\" />")
	assert.Contains(t, w.Body.String(), "invalid username or password")
	loginTestTeardown()
}

func TestLoginSuccess(t *testing.T) {
	loginTestSetup()

	(&models.User{ Name: "foo" }).CreateFromPassword("testpass")

	data := url.Values{ "username": {"foo"}, "password": {"testpass"} }
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/login",  bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	LoginPost(w, r)

	assert.Equal(t, "/profile", w.Header().Get("Location"))
	assert.Equal(t, 307, w.Code)

	loginTestTeardown()
}
