package nagiosplugin

import "fmt"

// Nagios plugin exit status.
type Status uint

// https://nagios-plugins.org/doc/guidelines.html#AEN78
const (
	OK Status = iota
	WARNING
	CRITICAL
	UNKNOWN
)

type (
	// Check results are ordered by severity: only the most severe check
	// result will be captured in the plugin's exit status. A status
	// policy is used to define severity as a function of check status.
	// Higher relative statusSeverity values assign higher severity to a
	// status. (Absolute values are insignificant.)
	statusSeverity uint
	statusPolicy   map[Status]statusSeverity
)

// NewDefaultStatusPolicy returns a status policy that assigns relative
// severity in accordance with conventional Nagios plugin return codes.
// Statuses associated with higher return codes are more severe.
func NewDefaultStatusPolicy() *statusPolicy {
	return &statusPolicy{
		OK:       statusSeverity(OK),
		WARNING:  statusSeverity(WARNING),
		CRITICAL: statusSeverity(CRITICAL),
		UNKNOWN:  statusSeverity(UNKNOWN),
	}
}

// NewStatusPolicy returns a status policy that assigns relative
// severity in accordance with a user-configurable prioritised slice.
// Check statuses must be listed in ascending severity order.
func NewStatusPolicy(statuses []Status) (*statusPolicy, error) {
	newPol := make(statusPolicy)
	for i, status := range statuses {
		newPol[status] = statusSeverity(i)
	}

	// Ensure all statuses are covered by the new policy.
	defaultPol := NewDefaultStatusPolicy()
	for status, _ := range *defaultPol {
		_, ok := newPol[status]
		if !ok {
			return nil, fmt.Errorf("missing status: %v", status)
		}
	}

	return &newPol, nil
}

// Returns string representation of a Status. Panics if given an invalid
// status (this will be recovered in check.Finish if it has been deferred).
func (s Status) String() string {
	switch s {
	case OK:
		return "OK"
	case WARNING:
		return "WARNING"
	case CRITICAL:
		return "CRITICAL"
	case UNKNOWN:
		return "UNKNOWN"
	}
	panic("Invalid nagiosplugin.Status.")
}

// Result encapsulates a machine-readable result code and a
// human-readable description of a problem. A check may have multiple
// Results. Only the most severe Result will be reported on the first
// line of plugin output and in the plugin's exit status.
type Result struct {
	status  Status
	message string
}
