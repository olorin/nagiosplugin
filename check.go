package nagiosplugin

import (
	"fmt"
	"os"
	"strings"
)

const (
	MessageSeparator = ", "
)

// Standalone Exit function for simple checks without multiple results
// or perfdata.
func Exit(status Status, message string) {
	fmt.Printf("%v: %s\n", status, message)
	os.Exit(int(status))
}

// Represents the state of a Nagios check.
type Check struct {
	results []Result
	status  Status
}

// AddResult adds a check result. This will not terminate the check. If
// status is the highest yet reported, this will update the check's
// final return status.
func (c Check) AddResult(status Status, message string) {
	var result Result
	result.status = status
	result.message = message
	c.results = append(c.results, result)
	if result.status > c.status {
		c.status = result.status
	}
}

// exitInfoText returns the most important result text, formatted for
// the first line of plugin output.
//
// Returns joined string of (MessageSeparator-separated) info text from
// results which have a status of at least c.status.
func (c Check) exitInfoText() string {
	importantMessages := make([]string, 0)
	for _, result := range c.results {
		if result.status >= c.status {
			importantMessages = append(importantMessages, result.message)
		}
	}
	return strings.Join(importantMessages, MessageSeparator)
}

// String representation of the check results, suitable for output and
// parsing by Nagios.
func (c Check) String() string {
	return fmt.Sprintf("%v: %s", c.status, c.exitInfoText())
}

// Finish ends the check, prints its output (to stdout), and exits with
// the correct status.
func (c Check) Finish() {
	fmt.Println(c)
	os.Exit(int(c.status))
}
