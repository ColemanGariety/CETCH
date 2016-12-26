package models

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)


func setup() {
	InitDB(os.Getenv("dbname"))
	UserCreate("foo", "bar")
}

func teardown() {
	UserDeleteAll()
	CloseDB()
}

func TestUserByName(t *testing.T) {
	setup()
	user, err := UserByName("foo")
	assert.Nil(t, err)
	assert.NotNil(t, user.Name)
	assert.NotNil(t, user.PasswordHash)
	teardown()
}

func TestUserExistsByName(t *testing.T) {
	setup()
	exists := UserExistsByName("foo")
	assert.Equal(t, exists, true)
	teardown()
}

func TestUserCreate(t *testing.T) {
	setup()
	UserDeleteAll()
	InitDB(os.Getenv("dbname"))
	err := UserCreate("foo", "bar")
	assert.Nil(t, err)
	teardown()
}

func TestUserDelete(t *testing.T) {
	setup()
	err := UserDelete("foo")
	assert.Nil(t, err)
	err = UserDelete("foo")
	assert.NotNil(t, err)
}
