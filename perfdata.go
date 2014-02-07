package nagiosplugin

import (
	"fmt"
	"math"
	"strings"
)

// PerfDatum represents one metric to be reported as part of a check
// result.
type PerfDatum struct {
	label string
	value float64
	unit  string
	min   *float64
	max   *float64
	warn  *float64
	crit  *float64
}

// validUnit returns true if the string is a valid UOM; otherwise false.
// It is case-insensitive.
func validUnit(unit string) bool {
	switch strings.ToLower(unit) {
	case "", "us", "ms", "s", "%", "b", "kb", "mb", "gb", "tb", "c":
		return true
	}
	return false
}

// NewPerfDatum returns a PerfDatum object suitable to use in a check
// result. unit must a valid Nagios unit, i.e., one of "us", "ms", "s",
// "%", "b", "kb", "mb", "gb", "tb", "c", or the empty string.
//
// Zero to four thresholds may be supplied: min, max, warn and crit.
// Thresholds may be positive or negative infinity, in which case they
// will be omitted in check output.
func NewPerfDatum(label string, unit string, value float64, thresholds ...float64) (*PerfDatum, error) {
	datum := new(PerfDatum)
	datum.label = label
	datum.value = value
	datum.unit = unit
	if !validUnit(unit) {
		return nil, fmt.Errorf("Invalid unit %v", unit)
	}
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return nil, fmt.Errorf("Perfdata value may not be infinity or NaN: %v.", value)
	}
	if len(thresholds) >= 1 {
		datum.min = &thresholds[0]
	}
	if len(thresholds) >= 2 {
		datum.max = &thresholds[1]
	}
	if len(thresholds) >= 3 {
		datum.warn = &thresholds[2]
	}
	if len(thresholds) >= 4 {
		datum.crit = &thresholds[3]
	}
	return datum, nil
}

// isThresholdSet returns true if one of min, max, warn or crit are set
// and false otherwise. They are determined to be 'set' if they are not
// a) the nil pointer or b) (either) infinity.
func isThresholdSet(t *float64) bool {
	if t == nil {
		return false
	}
	if math.IsInf(*t, 0) {
		return false
	}
	return true
}

// fmtThreshold returns a string representation of min, max, warn or
// crit (whether or not they are set).
func fmtThreshold(t *float64) string {
	if !isThresholdSet(t) {
		return ""
	}
	return fmt.Sprintf("%f", *t)
}

// String returns the string representation of a PerfDatum, suitable for
// check output.
func (p PerfDatum) String() string {
	value := fmt.Sprintf("%s=%f%s", p.label, p.value, p.unit)
	value += fmt.Sprintf(";%s;%s", fmtThreshold(p.warn), fmtThreshold(p.crit))
	value += fmt.Sprintf(";%s;%s", fmtThreshold(p.min), fmtThreshold(p.max))
	return value
}

// RenderPerfdata accepts a slice of PerfDatum objects and returns their
// concatenated string representations in a form suitable to append to
// the first line of check output.
func RenderPerfdata(perfdata []PerfDatum) string {
	value := ""
	if len(perfdata) == 0 {
		return value
	}
	// Demarcate start of perfdata in check output.
	value += " |"
	for _, datum := range perfdata {
		value += fmt.Sprintf(" %v", datum)
	}
	return value
}
