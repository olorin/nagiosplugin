[![Build status](https://travis-ci.org/olorin/nagiosplugin.svg?branch=master)](https://travis-ci.org/olorin/nagiosplugin)

# nagiosplugin

Package for writing Nagios/Icinga/et cetera plugins in Go (golang).

# Documentation

See http://godoc.org/github.com/olorin/nagiosplugin. 

# Usage example

The general usage pattern looks like this:

```go
func main() {
	// Initialize the check - this will return an UNKNOWN result
	// until more results are added.
	check := nagiosplugin.NewCheck()
	// If we exit early or panic() we'll still output a result.
	defer check.Finish()

	// obtain data here

	// Add an 'OK' result - if no 'worse' check results have been
	// added, this is the one that will be output.
	check.AddResult(nagiosplugin.OK, "everything looks shiny, cap'n")
	// Add some perfdata too (label, unit, value, min, max,
	// warn, crit). The math.Inf(1) will be parsed as 'no
	// maximum'.
	check.AddPerfDatum("badness", "kb", 3.14159, 0.0, math.Inf(1), 8000.0, 9000.0)

	// Parse a range from the command line and warn on a match.
	warnRange, err := nagiosplugin.ParseRange( "1:2" )
	if err != nil {
		check.AddResult(nagiosplugin.UNKNOWN, "error parsing warning range")
	}
	if warnRange.Check( 3.14159 ) {
		check.AddResult(nagiosplugin.WARNING, "Are we crashing again?")
	}
}
```

In the example above, multiple results were added to the check with `AddResult()`. The final state of the check will be determined by its most severe result. By default, result severity is ordered to match the plugin return codes documented in the [Nagios plugin developer guidelines][guidelines]. A WARNING result is considered more severe than OK, CRITICAL more severe than WARNING, and UNKNOWN most severe of all. A status policy may be used to modify this default behaviour:

```go
func main() {
	check := nagiosplugin.NewCheckWithOptions(nagiosplugin.CheckOptions{
		// OK -> UNKNOWN -> WARNING -> CRITICAL
		StatusPolicy: nagiosplugin.NewOUWCStatusPolicy(),
	})
	defer check.Finish()

	// Now, any WARNING or CRITICAL results added to the check will be
	// considered more severe than any UNKNOWN result.
}
```

[guidelines]: https://nagios-plugins.org/doc/guidelines.html

# Language version

Requires go >= 1.0; tested with versions up to 1.6.
