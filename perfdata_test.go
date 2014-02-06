package nagiosplugin

import (
	"testing"
	"math"
)

func TestPerfdata(t *testing.T) {
	expected := "badness=9003.400000ms;4000.000000;9000.000000;10.000000;"
	pd, err := NewPerfDatum("badness", "ms", 9003.4, 10.0, math.Inf(1), 4000.0, 9000.0)
	if err != nil {
		t.Errorf("Could not render perfdata: %v", err)
	}
	if pd.String() != expected {
		t.Errorf("Perfdata rendering error: expected %s, got %v", expected, pd)
	}
}
