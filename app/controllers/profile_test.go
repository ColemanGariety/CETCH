package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestProfileTemporaryRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/profile", http.HandlerFunc(ProfileShow))
	mux.Handle("/login", http.HandlerFunc(LoginShow))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/profile")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get(u.String())

	assert.NoError(t, err)
	assert.Equal(t, 307, res.StatusCode)
}
