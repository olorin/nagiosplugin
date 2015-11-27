package nagiosplugin

import (
	"testing"
)

// see https://www.monitoring-plugins.org/doc/guidelines.html#THRESHOLDFORMAT

// TestInsideRange ensures that an explicitly bounded range only accepts
// values within the range.
func TestInsideRange(t *testing.T) {
	var r *Range
	var err error

	rangeStr := "10:20"
	r, err = ParseRange(rangeStr)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", rangeStr, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-1.0, true},
		{0.0, true},
		{1.0, true},
		{9.0, true},
		{10.0, false},
		{15.0, false},
		{20.0, false},
		{21.0, true},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be %v", test.value, test.shouldAlert)
		}
	}
}

// TestOutsideRange ensures that an explicitly bounded range prefixed
// with the at sign (@) only accepts values outside the range.
func TestOutsideRange(t *testing.T) {
	var r *Range
	var err error

	rangeStr := "@10:20"
	r, err = ParseRange(rangeStr)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", rangeStr, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-1.0, false},
		{0.0, false},
		{1.0, false},
		{9.0, false},
		{10.0, true},
		{15.0, true},
		{20.0, true},
		{21.0, false},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be %v", test.value, test.shouldAlert)
		}
	}
}

// TestImpliedMinimumRange ensures that a range string with no explicit
// minimum defaults to a minimum of 0.
func TestImpliedMinimumRange(t *testing.T) {
	var r *Range
	var err error

	rangeStr := "10"
	r, err = ParseRange(rangeStr)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", rangeStr, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-1.0, true},
		{0.0, false},
		{5.0, false},
		{11.0, true},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be %v", test.value, test.shouldAlert)
		}
	}
}

// TestImpliedMaximumRange ensures that a range string with no explicit
// maximum defaults to a maximum of +Inf.
func TestImpliedMaximumRange(t *testing.T) {
	var r *Range
	var err error

	rangeStr := "10:"
	r, err = ParseRange(rangeStr)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", rangeStr, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-1.0, true},
		{0.0, true},
		{1.0, true},
		{5.0, true},
		{10.0, false},
		{11.0, false},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be %v", test.value, test.shouldAlert)
		}
	}
}

// TestNegInfRange ensures that a range string with a minimum bound of a
// tilde is correctly interpreted to mean -Inf.
func TestNegInfRange(t *testing.T) {
	var r *Range
	var err error

	rangeStr := "~:10"
	r, err = ParseRange(rangeStr)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", rangeStr, err)
	}

	tests := []struct {
		value       float64
		shouldAlert bool
	}{
		{-1.0, false},
		{0.0, false},
		{1.0, false},
		{5.0, false},
		{10.0, false},
		{11.0, true},
	}
	for _, test := range tests {
		didAlert := r.Check(test.value)
		if didAlert != test.shouldAlert {
			t.Errorf("Check(%v) should be %v", test.value, test.shouldAlert)
		}
	}
}

// TestInvalidRanges ensures that ParseRange correctly fails on input
// that is known to be invalid.
func TestInvalidRanges(t *testing.T) {
	var err error

	badRanges := []string{
		"20:10", // Violates min <= max
		"10,20", // The comma is non-sensical
	}
	for _, rangeStr := range badRanges {
		_, err = ParseRange(rangeStr)
		if err == nil {
			t.Errorf("ParseRange(%v) should have returned an error", rangeStr)
		}
	}
}
