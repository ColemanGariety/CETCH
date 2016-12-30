package utils

import (
	"strings"
	"unicode"
	"time"
)

func StripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func NextFriday() time.Time {
	date := time.Now().AddDate(0, 0, 1)
	for {
		if date.Weekday() == time.Friday {
			date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
			break
		}
		date = date.AddDate(0, 0, 1)
	}
	return date
}
