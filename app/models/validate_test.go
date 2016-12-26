package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type MiscForm struct {
	Form
	Foo string
	Bar string
}

func TestValidatePresence(t *testing.T) {
	form := &MiscForm{
		Foo: "",
		Bar: "",
	}

	form.Errors = make(map[string]string)
	form.ValidatePresence(form.Foo, "Foo")
	form.ValidatePresence(form.Bar, "Bar")
	
	expected := 2
	actual := len(form.Errors)
	
	assert.Equal(t, expected, actual, "")
}

func TestValidateNoSpace(t *testing.T) {
	form := &MiscForm{
		Foo: "f o o",
	}

	form.Errors = make(map[string]string)
	form.ValidateNoSpace(form.Foo, "Foo")
	
	expected := 1
	actual := len(form.Errors)

	assert.Equal(t, expected, actual, "")
}

func TestValidateConfirmation(t *testing.T) {
	form := &MiscForm{
		Foo: "foo",
		Bar: "bar",
	}

	form.Errors = make(map[string]string)
	form.ValidateConfirmation(form.Foo, "Foo", form.Bar, "Bar")
	
	expected := 1
	actual := len(form.Errors)

	assert.Equal(t, expected, actual, "")
}

func TestValidateLength(t *testing.T) {
	form := &MiscForm{
		Foo: "2shrt",
		Bar: "toooloooonnnnggggsttttrrrriiiinnngggg",
	}

	form.Errors = make(map[string]string)
	form.ValidateLength(form.Foo, "Foo", 6, 10)
	form.ValidateLength(form.Bar, "Bar", 6, 10)
	
	expected := 2
	actual := len(form.Errors)

	assert.Equal(t, expected, actual, "")
}

func TestFieldIsValid(t *testing.T) {
	form := &MiscForm{
		Foo: "foo",
	}

	form.Errors = make(map[string]string)

	form.ValidatePresence(form.Foo, "Foo")
	form.ValidateNoSpace(form.Foo, "Foo")
	form.ValidateLength(form.Foo, "Foo", 3, 3)
	form.ValidateConfirmation(form.Foo, "Foo", form.Foo, "Foo")
	
	expected := true
	actual := form.FieldIsValid("Foo")

	assert.Equal(t, expected, actual, "")
}

func TestIsValid(t *testing.T) {
	form := &MiscForm{
		Foo: "foo",
	}

	form.Errors = make(map[string]string)

	form.ValidatePresence(form.Foo, "Foo")
	form.ValidateNoSpace(form.Foo, "Foo")
	form.ValidateLength(form.Foo, "Foo", 3, 3)
	form.ValidateConfirmation(form.Foo, "Foo", form.Foo, "Foo")

	expected := true
	actual := form.IsValid()

	assert.Equal(t, expected, actual, "")
}

func TestIsNotValid(t *testing.T) {
	form := &MiscForm{
		Foo: "",
	}

	form.Errors = make(map[string]string)

	form.ValidatePresence(form.Foo, "Foo")
	form.ValidateNoSpace(form.Foo, "Foo")
	form.ValidateLength(form.Foo, "Foo", 3, 3)
	form.ValidateConfirmation(form.Foo, "Foo", form.Foo, "Foo")

	expected := false
	actual := form.IsValid()

	assert.Equal(t, expected, actual, "")
}
