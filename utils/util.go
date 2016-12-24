package utils

import (
	"strings"
	"unicode"
)

func StripSpace(str string) (string) {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
