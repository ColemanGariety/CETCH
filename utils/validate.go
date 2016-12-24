package utils

import (
	"fmt"
	"strings"
	"log"
)

type Form struct {
	Errors map[string]string
}

func (form Form) ValidatePresence(value string, field string) bool {
	if strings.TrimSpace(value) == "" {
		form.Errors[field] = fmt.Sprintf("%s must not be blank", field)
		return false
	}
	return true
}

func (form Form) ValidateNoSpace(value string, field string) bool {
	log.Println(strings.TrimSpace(value))
	if StripSpace(value) != value {
		form.SetError(field, fmt.Sprintf("%s may not contain spaces", field))
		return false
	}
	return true
}

func (form Form) ValidateConfirmation(value string, field string, confirmationValue string, confirmationField string) bool {
	if value != confirmationValue {
		form.Errors[confirmationField] = fmt.Sprintf("%s and %s must match", field, confirmationField)
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
