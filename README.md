nagiosplugin
============

Package for writing Nagios/Icinga/et cetera plugins in Go (golang).

This package is a quick hack/work in progress (contributions welcome). In
particular, it currently lacks

 - threshold parsing
 - in-check error handling (returning the correct status if the plugin
   dies mid-run)

documentation
=============

See http://godoc.org/github.com/fractalcat/nagiosplugin. 
