package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidatePresence(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "",
		"bar": "",
	}

	props.ValidatePresence("foo")
	props.ValidatePresence("bar")

	expected := 2
	actual := len(props["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateNoSpace(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "f o o",
	}

	props.ValidateNoSpace("foo")

	expected := 1
	actual := len(props["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateConfirmation(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "foo",
		"bar": "bar",
	}

	props.ValidateConfirmation("foo", "bar")

	expected := 1
	actual := len(props["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateLength(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "2shrt",
		"bar": "toooloooonnnnggggsttttrrrriiiinnngggg",
	}

	props.ValidateLength("foo", 6, 10)
	props.ValidateLength("bar", 6, 10)

	expected := 2
	actual := len(props["errors"].(map[string]string))

	assert.Equal(t, expected, actual, "")
}

func TestValidateEmail(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "foobar.com",
	}

	props.ValidateEmail("foo")
	assert.Equal(t, 1, len(props["errors"].(map[string]string)))
}

func TestFieldIsValid(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "foo",
	}

	props.ValidatePresence("foo")
	props.ValidateNoSpace("foo")
	props.ValidateLength("foo", 3, 3)
	props.ValidateConfirmation("foo", "foo")

	expected := true
	actual := props.FieldIsValid("foo")

	assert.Equal(t, expected, actual, "")
}

func TestIsValid(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "foo",
	}

	props.ValidatePresence("foo")
	props.ValidateNoSpace("foo")
	props.ValidateLength("foo", 3, 3)
	props.ValidateConfirmation("foo", "foo")

	expected := true
	actual := props.IsValid()

	assert.Equal(t, expected, actual, "")
}

func TestIsNotValid(t *testing.T) {
	props := Props{
		"errors": make(map[string]string),
		"foo": "",
	}

	props.ValidatePresence("foo")
	props.ValidateNoSpace("foo")
	props.ValidateLength("foo", 3, 3)
	props.ValidateConfirmation("foo", "foo")

	expected := false
	actual := props.IsValid()

	assert.Equal(t, expected, actual, "")
}
