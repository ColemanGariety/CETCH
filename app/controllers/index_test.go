package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"bytes"
	"os"

	"github.com/JacksonGariety/cetch/app/models"
)

func setup() {
	models.InitDB(os.Getenv("dbname"))
}

func teardown() {
	models.CloseDB()
}

func TestIndexOK(t *testing.T) {
	setup()
	ts := httptest.NewServer(http.HandlerFunc(Index))
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/")

	res, err := http.Get(u.String())

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	teardown()
}
