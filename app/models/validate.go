package models

import (
	"fmt"
	"strings"

	"github.com/JacksonGariety/cetch/app/utils"
)

type Form map[string]interface{}

func (form Form) ValidatePresence(field string) bool {
	if strings.TrimSpace(form[field].(string)) == "" {
		form.SetError(field, fmt.Sprintf("%s can't be blank", field))
		return false
	}
	return true
}

func (form Form) ValidateNoSpace(field string) bool {
	if utils.StripSpaces(form[field].(string)) != form[field] {
		form.SetError(field, fmt.Sprintf("%s may not contain spaces", field))
		return false
	}
	return true
}

func (form Form) ValidateConfirmation(field string, confirmationField string) bool {
	if form[field] != form[confirmationField] {
		form.SetError(confirmationField, fmt.Sprintf("%s and %s must match", field, confirmationField))
		return false
	}
	return true
}

func (form Form) ValidateEmail(field string) bool {
	if !(strings.Contains(form[field].(string), "@")) {
		form.SetError(field, fmt.Sprintf("%s must be an email", field))
		return false
	}
	return true
}

func (form Form) ValidateLength(field string, min int, max int) bool {
	length := len(form[field].(string))
	if length < min || length > max {
		form.SetError(field, fmt.Sprintf("%s must be between %d and %d characters in length", field, min, max))
		return false
	}
	return true
}

func (form Form) FieldIsValid(field string) bool {
	return form["errors"].(map[string]string)[field] == ""
}

func (form Form) IsValid() bool {
	return len(form["errors"].(map[string]string)) == 0
}

func (form Form) SetError(field string, value string) {
	form["errors"].(map[string]string)[field] = value
}
