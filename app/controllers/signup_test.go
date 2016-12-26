package controllers

import (
	"testing"
	"net/http"
	"os"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"bytes"
	"net/url"

	"github.com/JacksonGariety/cetch/app/models"
)

func signupTestSetup() {
	models.InitDB(os.Getenv("dbname"))
}

func signupTestTeardown() {
	(&models.Users{}).DeleteAll()
	models.CloseDB()
}

func TestSignupShowOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SignupShow))
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/signup")

	res, err := http.Get(u.String())

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

func TestSignupInUseUsername(t *testing.T) {
	signupTestSetup()

	// make the user
	(&models.User{ Name: "foo" }).CreateFromPassword("bar")

	data := url.Values{ "email": {"foo@bar.raz"}, "username": {"foo"}, "password": {"bar"}, "password_confirmation": {"bar"} }
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/signup",  bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	SignupPost(w, r)

	assert.Contains(t, w.Body.String(), "<input type=\"email\" name=\"email\" value=\"foo@bar.raz\" />")
	assert.Contains(t, w.Body.String(), "<input type=\"text\" name=\"username\" value=\"foo\" />")
	assert.Contains(t, w.Body.String(), "<input type=\"password\" name=\"password\" value=\"bar\" />")
	assert.Contains(t, w.Body.String(), "<input type=\"password\" name=\"password_confirmation\" value=\"bar\" />")
	assert.Contains(t, w.Body.String(), "username is already in use")
	signupTestTeardown()
}

func TestSignupSuccess(t *testing.T) {
	signupTestSetup()

	data := url.Values{ "email": {"foo@bar.raz"}, "username": {"foo"}, "password": {"testpass"}, "password_confirmation": {"testpass"} }
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/signup",  bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	SignupPost(w, r)

	assert.Equal(t, "/profile", w.Header().Get("Location"))
	assert.Equal(t, 307, w.Code)

	signupTestTeardown()
}
