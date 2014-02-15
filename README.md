nagiosplugin
============

Package for writing Nagios/Icinga/et cetera plugins in Go (golang).

documentation
=============

See http://godoc.org/github.com/fractalcat/nagiosplugin. 

usage example
=============

The general usage pattern looks like this:

	func main() {
		// Initialize the check - this will return an UNKNOWN result
		// until more perfdata are added.
		check := nagiosplugin.NewCheck()
		// If we exit early or panic() we'll still output a result.
		defer check.Finish()
	
		// obtain data here
	
		// Add an 'OK' result - if no 'worse' check results have been
		// added, this is the one that will be output.
		check.AddResult(nagiosplugin.OK, "everything looks shiny, cap'n")
		// Add some perfdata too (label, unit, min, max, warn, crit).
		// The math.Inf(1) will be parsed as 'no maximum'. 
		check.AddPerfDatum("badness", "kb", 0.0, math.Inf(1), 8000.0, 9000.0)
	}
