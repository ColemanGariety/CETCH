package controllers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func setup() {
	utils.InitTemplates()
	models.InitDB("user=cetch dbname=cetch_test sslmode=disable")
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
