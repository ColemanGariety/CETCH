package models

import (
	"fmt"
	"strings"

	"github.com/JacksonGariety/cetch/app/utils"
)

type Form struct {
	Errors map[string]string
}

func (form Form) ValidatePresence(value string, field string) bool {
	if strings.TrimSpace(value) == "" {
		form.SetError(field, fmt.Sprintf("%s must not be blank", field))
		return false
	}
	return true
}

func (form Form) ValidateNoSpace(value string, field string) bool {
	if utils.StripSpaces(value) != value {
		form.SetError(field, fmt.Sprintf("%s may not contain spaces", field))
		return false
	}
	return true
}

func (form Form) ValidateConfirmation(value string, field string, confirmationValue string, confirmationField string) bool {
	if value != confirmationValue {
		form.SetError(confirmationField, fmt.Sprintf("%s and %s must match", field, confirmationField))
		return false
	}
	return true
}

func (form Form) ValidateLength(value string, field string, min int, max int) bool {
	length := len(value)
	if length < min || length > max {
		form.SetError(field, fmt.Sprintf("%s must be between %d and %d characters in length", field, min, max))
		return false
	}
	return true
}

func (form Form) FieldIsValid(field string) bool {
	return form.Errors[field] == ""
}

func (form Form) IsValid() bool {
	return len(form.Errors) == 0
}

func (form Form) SetError(field string, value string) {
	form.Errors[field] = value
}
