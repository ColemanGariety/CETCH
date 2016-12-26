package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClaimsCreate(t *testing.T) {
	signedToken, _ := ClaimsCreate("foo")
	assert.NotNil(t, signedToken)
}
