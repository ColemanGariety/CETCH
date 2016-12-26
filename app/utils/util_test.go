package utils

import (
	"testing"
)

func TestStripSpaces(t *testing.T) {
	expected := "foobar"
	actual := StripSpaces(" foo bar ")
	if expected != actual {
		t.Error("expected %s to equal %s", expected, actual)
	}
}
