package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidatePresence(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "",
		"bar": "",
	}

	form.ValidatePresence("foo")
	form.ValidatePresence("bar")

	expected := 2
	actual := len(form["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateNoSpace(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "f o o",
	}

	form.ValidateNoSpace("foo")

	expected := 1
	actual := len(form["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateConfirmation(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "foo",
		"bar": "bar",
	}

	form.ValidateConfirmation("foo", "bar")

	expected := 1
	actual := len(form["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateLength(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "2shrt",
		"bar": "toooloooonnnnggggsttttrrrriiiinnngggg",
	}

	form.ValidateLength("foo", 6, 10)
	form.ValidateLength("bar", 6, 10)

	expected := 2
	actual := len(form["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateEmail(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "foobar.com",
	}

	form.ValidateEmail("foo")
	assert.Equal(t, 1, len(form["errors"].(map[string]string)))
}

func TestFieldIsValid(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "foo",
	}

	form.ValidatePresence("foo")
	form.ValidateNoSpace("foo")
	form.ValidateLength("foo", 3, 3)
	form.ValidateConfirmation("foo", "foo")

	expected := true
	actual := form.FieldIsValid("foo")

	assert.Equal(t, expected, actual, "")
}

func TestIsValid(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "foo",
	}

	form.ValidatePresence("foo")
	form.ValidateNoSpace("foo")
	form.ValidateLength("foo", 3, 3)
	form.ValidateConfirmation("foo", "foo")

	expected := true
	actual := form.IsValid()

	assert.Equal(t, expected, actual, "")
}

func TestIsNotValid(t *testing.T) {
	form := Form{
		"errors": make(map[string]string),
		"foo": "",
	}

	form.ValidatePresence("foo")
	form.ValidateNoSpace("foo")
	form.ValidateLength("foo", 3, 3)
	form.ValidateConfirmation("foo", "foo")

	expected := false
	actual := form.IsValid()

	assert.Equal(t, expected, actual, "")
}
