package nagiosplugin

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Range is a combination of a lower boundary, an upper boundary
// and a flag for inverted (@) range semantics.
type Range struct {
	Start         float64
	End           float64
	AlertOnInside bool
}

// Returns a new range object and nil if the given
// range definition was valid or nil and an error.
func ParseRange(rangeStr string) (*Range, error) {
	// Set Defaults
	t := &Range{
		Start:         0,
		End:           math.Inf(1),
		AlertOnInside: false,
	}
	// Remove leading and trailing whitespaces
	rangeStr = strings.Trim(rangeStr, " \n\r")

	// Check for inverted semantics
	if rangeStr[0] == '@' {
		t.AlertOnInside = true
		rangeStr = rangeStr[1:]
	}

	// parse lower limit
	endPos := strings.Index(rangeStr, ":")
	if endPos > -1 {
		if rangeStr[0] == '~' {
			t.Start = math.Inf(-1)
		} else {
			min, err := strconv.ParseFloat(rangeStr[0:endPos], 64)
			if err != nil {
				fmt.Println("failed to parse lower limit:", err)
			} else {
				t.Start = min
			}
		}
		rangeStr = rangeStr[endPos+1:]
	}

	// parse upper limit
	if len(rangeStr) > 0 {
		max, err := strconv.ParseFloat(rangeStr, 64)
		if err != nil {
			fmt.Println("failed to parse upper limit:", err)
		} else {
			t.End = max
		}
	}

	if t.End < t.Start {
		return nil, errors.New("Invalid range definition. min <= max violated!")
	}

	// OK
	return t, nil
}

// Returns true if an alert should be raised based on the range.
func (r *Range) Check(value float64) bool {
	var no bool = false
	var yes bool = true

	if r.AlertOnInside {
		no = true
		yes = false
	}
	// see https://www.monitoring-plugins.org/doc/guidelines.html#THRESHOLDFORMAT
	if r.Start <= value && value <= r.End {
		return no
	} else {
		return yes
	}
	return yes
}

// CheckInt is a convenience method which does an unchecked type
// conversion from an int to a float64.
func (r *Range) CheckInt(val int) bool {
	return r.Check(float64(val))
}

// CheckUint64 is a convenience method which does an unchecked type
// conversion from an uint64 to a float64.
func (r *Range) CheckUint64(val uint64) bool {
	return r.Check(float64(val))
}
