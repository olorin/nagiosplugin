package nagiosplugin_test

import (
	"github.com/olorin/nagiosplugin"
	"math"
)

func ExampleExit() {
	nagiosplugin.Exit(nagiosplugin.CRITICAL, "Badness over 9000!")
}

func Example() {
	check := nagiosplugin.NewCheck()
	// Make sure the check always (as much as possible) exits with
	// the correct output and return code if we terminate unexpectedly.
	defer check.Finish()
	// (If the check panicked on the next line, it'd exit with a
	// default UNKNOWN result.)
	//
	// Our check is testing the internal consistency of the
	// universe.
	value := math.Pi
	// We add a dimensionless metric with a minimum of zero, an
	// unbounded maximum, a warning threshold of 4000.0 and a
	// critical threshold of 9000.0 (for graphing purposes).
	check.AddPerfDatum("badness", "", value, 0.0, math.Inf(1), 4000.0, 9000.0)
	// Add an OK check result as the universe appears sane.
	check.AddResult(nagiosplugin.OK, "Everything looks shiny from here, cap'n")
	// We potentially perform more checks and add more results here;
	// if there's more than one, the highest result will be the one
	// returned (in ascending order OK, WARNING, CRITICAL, UNKNOWN).

	// This will print:
	// OK: Everything looks shiny from here, cap'n | badness=3.141592653589793;4000;9000;0;
}
