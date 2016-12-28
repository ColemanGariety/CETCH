package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClaimsCreate(t *testing.T) {
	signedToken, _, _ := ClaimsCreate("foo")
	assert.NotNil(t, signedToken)
}
