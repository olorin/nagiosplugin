package nagiosplugin

import (
	"math"
	"testing"
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

func TestRenderPerfdata(t *testing.T) {
	expected := " | goodness=3.141593kb;;;3.000000;34.557519 goodness=6.283185kb;;;3.000000;34.557519 goodness=9.424778kb;;;3.000000;34.557519 goodness=12.566371kb;;;3.000000;34.557519 goodness=15.707963kb;;;3.000000;34.557519 goodness=18.849556kb;;;3.000000;34.557519 goodness=21.991149kb;;;3.000000;34.557519 goodness=25.132741kb;;;3.000000;34.557519 goodness=28.274334kb;;;3.000000;34.557519 goodness=31.415927kb;;;3.000000;34.557519"
	pd := make([]PerfDatum, 0)
	for i := 0; i < 10; i++ {
		datum, err := NewPerfDatum("goodness", "kb", math.Pi * float64(i+1), 3.0, math.Pi * 11)
		if err != nil {
			t.Errorf("Could not create perfdata: %v", err)
		}
		pd = append(pd, *datum)
	}
	result := RenderPerfdata(pd)
	if result != expected {
		t.Errorf("Perfdata rendering error: expected %s, got %v", expected, result)
	}
}
