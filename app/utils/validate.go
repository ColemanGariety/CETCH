package utils

import (
	"fmt"
	"strings"
)

type Props map[string]interface{}

func (props Props) ValidatePresence(field string) bool {
	switch value := props[field].(type) {
	case string:
		if value == "" {
			props.SetError(field, fmt.Sprintf("%s can't be blank", field))
			return false
		}
	case float64:
		if value == 0.0 {
			props.SetError(field, fmt.Sprintf("%s can't be blank", field))

		}
	}
	return true
}

func (props Props) ValidateNoSpace(field string) bool {
	if StripSpaces(props[field].(string)) != props[field] {
		props.SetError(field, fmt.Sprintf("%s may not contain spaces", field))
		return false
	}
	return true
}

func (props Props) ValidateConfirmation(field string, confirmationField string) bool {
	if props[field] != props[confirmationField] {
		props.SetError(confirmationField, fmt.Sprintf("%s and %s must match", field, confirmationField))
		return false
	}
	return true
}

func (props Props) ValidateEmail(field string) bool {
	if !(strings.Contains(props[field].(string), "@")) {
		props.SetError(field, fmt.Sprintf("%s must be an email", field))
		return false
	}
	return true
}

func (props Props) ValidateLength(field string, min int, max int) bool {
	length := len(props[field].(string))
	if length < min || length > max {
		props.SetError(field, fmt.Sprintf("%s must be between %d and %d characters in length", field, min, max))
		return false
	}
	return true
}

func (props Props) FieldIsValid(field string) bool {
	return props["errors"].(map[string]string)[field] == ""
}

func (props Props) IsValid() bool {
	return len(props["errors"].(map[string]string)) == 0
}

func (props Props) SetError(field string, value string) {
	props["errors"].(map[string]string)[field] = value
}
