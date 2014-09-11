package nagiosplugin

import (
	"testing"
)

// see https://www.monitoring-plugins.org/doc/guidelines.html#THRESHOLDFORMAT

func TestRange(t *testing.T) {
	var r *Range
	var err error
	var pat string

	// true   => raise alert
	// false  => ok (no alert)

	/*
	 * Case 1: Pattern "10" -> Range {0..10}
	 * Tests:
	 * -1   -> true
	 *  0   -> false
	 *  5   -> false
	 * 10   -> false
	 * 11   -> true
	 */
	pat = "10"
	r, err = ParseRange(pat)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", pat, err)
	}
	if !r.Check(-1.0) {
		t.Errorf("Check -1.0 should be true")
	}
	if r.Check(0.0) {
		t.Errorf("Check 0.0 should be false")
	}
	if r.Check(5.0) {
		t.Errorf("Check 5.0 should be false")
	}
	if r.Check(10.0) {
		t.Errorf("Check 10.0 should be false")
	}
	if !r.Check(11.0) {
		t.Errorf("Check 11.0 should be true")
	}
	/*
	 * Case 2: Pattern: "10:" -> Range {10..inf}
	 * Tests:
	 * -1   -> true
	 *  0   -> true
	 *  1   -> true
	 *  5   -> true
	 * 10   -> false
	 * 11   -> false
	 */
	pat = "10:"
	r, err = ParseRange(pat)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", pat, err)
	}
	if !r.Check(-1.0) {
		t.Errorf("Check -1.0 should be true")
	}
	if !r.Check(0.0) {
		t.Errorf("Check 0.0 should be true")
	}
	if !r.Check(1.0) {
		t.Errorf("Check 1.0 should be true")
	}
	if !r.Check(5.0) {
		t.Errorf("Check 10.0 should be true")
	}
	if r.Check(10.0) {
		t.Errorf("Check 10.0 should be false")
	}
	if r.Check(11.0) {
		t.Errorf("Check 11.0 should be false")
	}
	/*
	 * Case 3: Pattern: "~:10" -> Range {-inf..10}
	 * Tests:
	 * -1   -> false
	 *  0   -> false
	 *  1   -> false
	 *  5   -> false
	 * 10   -> false
	 * 11   -> true
	 */
	pat = "~:10"
	r, err = ParseRange(pat)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", pat, err)
	}
	if r.Check(-1.0) {
		t.Errorf("Check -1.0 should be false")
	}
	if r.Check(0.0) {
		t.Errorf("Check 0.0 should be false")
	}
	if r.Check(1.0) {
		t.Errorf("Check 1.0 should be false")
	}
	if r.Check(5.0) {
		t.Errorf("Check 10.0 should be false")
	}
	if r.Check(10.0) {
		t.Errorf("Check 10.0 should be false")
	}
	if !r.Check(11.0) {
		t.Errorf("Check 11.0 should be true")
	}
	/*
	 * Case 4: Pattern: "10:20" -> Range {10..20}
	 * Tests:
	 * -1   -> true
	 *  0   -> true
	 *  1   -> true
	 *  9   -> true
	 * 10   -> false
	 * 15   -> false
	 * 20   -> false
	 * 21   -> true
	 */
	pat = "10:20"
	r, err = ParseRange(pat)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", pat, err)
	}
	if !r.Check(-1.0) {
		t.Errorf("Check -1.0 should be true")
	}
	if !r.Check(0.0) {
		t.Errorf("Check 0.0 should be true")
	}
	if !r.Check(1.0) {
		t.Errorf("Check 1.0 should be true")
	}
	if !r.Check(9.0) {
		t.Errorf("Check 9.0 should be true")
	}
	if r.Check(10.0) {
		t.Errorf("Check 10.0 should be false")
	}
	if r.Check(15.0) {
		t.Errorf("Check 15.0 should be false")
	}
	if r.Check(20.0) {
		t.Errorf("Check 20.0 should be false")
	}
	if !r.Check(21.0) {
		t.Errorf("Check 21.0 should be true")
	}
	/*
	 * Case 5: Pattern: "@10:20" -> Range !{10..20}
	 * Tests:
	 * -1   -> false
	 *  0   -> false
	 *  1   -> false
	 *  9   -> false
	 * 10   -> true
	 * 15   -> true
	 * 20   -> true
	 * 21   -> false
	 */
	pat = "@10:20"
	r, err = ParseRange(pat)
	if err != nil {
		t.Fatalf("Failed to parse %v: %v", pat, err)
	}
	if r.Check(-1.0) {
		t.Errorf("Check -1.0 should be false")
	}
	if r.Check(0.0) {
		t.Errorf("Check 0.0 should be false")
	}
	if r.Check(1.0) {
		t.Errorf("Check 1.0 should be false")
	}
	if r.Check(9.0) {
		t.Errorf("Check 9.0 should be false")
	}
	if !r.Check(10.0) {
		t.Errorf("Check 10.0 should be true")
	}
	if !r.Check(15.0) {
		t.Errorf("Check 15.0 should be true")
	}
	if !r.Check(20.0) {
		t.Errorf("Check 20.0 should be true")
	}
	if r.Check(21.0) {
		t.Errorf("Check 21.0 should be false")
	}
}
