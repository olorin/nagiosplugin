package nagiosplugin

// Nagios plugin exit status.
type Status uint

const (
	OK Status = iota
	WARNING
	CRITICAL
	UNKNOWN
)

// Returns string representation of a Status. Panics if given an invalid
// status. 
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

type Result struct {
	status Status
	message string
}

