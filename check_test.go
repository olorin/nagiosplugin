package nagiosplugin

import (
	"testing"
)

func TestDefaultCheck(t *testing.T) {
	c := NewDefaultCheck()
	expected := "UNKNOWN: No check results specified"
	result := c.String()
	if expected != result {
		t.Errorf("Expected check output %v, got check output %v", expected, result)
	}
}

func TestCheck(t *testing.T) {
	c := NewCheck(OK, "everything looks shiny from here, cap'n")
	expected := "OK: everything looks shiny from here, cap'n"
	result := c.String()
	if expected != result {
		t.Errorf("Expected check output %v, got check output %v", expected, result)
	}
}
