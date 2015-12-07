package nagiosplugin

import (
	"math"
	"testing"
)

func TestPerfdata(t *testing.T) {
	expected := "badness=9003.4ms;4000;9000;10;"
	pd, err := NewPerfDatum("badness", "ms", 9003.4, 10.0, math.Inf(1), 4000.0, 9000.0)
	if err != nil {
		t.Errorf("Could not render perfdata: %v", err)
	}
	if pd.String() != expected {
		t.Errorf("Perfdata rendering error: expected %s, got %v", expected, pd)
	}
}

func TestRenderPerfdata(t *testing.T) {
	expected := " | goodness=3.141592653589793kb;;;3;34.55751918948773 goodness=6.283185307179586kb;;;3;34.55751918948773 goodness=9.42477796076938kb;;;3;34.55751918948773 goodness=12.566370614359172kb;;;3;34.55751918948773 goodness=15.707963267948966kb;;;3;34.55751918948773 goodness=18.84955592153876kb;;;3;34.55751918948773 goodness=21.991148575128552kb;;;3;34.55751918948773 goodness=25.132741228718345kb;;;3;34.55751918948773 goodness=28.274333882308138kb;;;3;34.55751918948773 goodness=31.41592653589793kb;;;3;34.55751918948773"
	pd := make([]PerfDatum, 0)
	for i := 0; i < 10; i++ {
		datum, err := NewPerfDatum("goodness", "kb", math.Pi*float64(i+1), 3.0, math.Pi*11)
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

func TestRenderPerfdataWithOmissions(t *testing.T) {
	pd := make([]PerfDatum, 0)
	datum, err := NewPerfDatum(
		"age",       // label
		"s",         // UOM
		0.123,       // value
		0.0,         // min
		math.Inf(1), // max: +Inf -> omit
		math.NaN(),  // warn: NaN -> omit
		0.5)         // crit
	if err != nil {
		t.Errorf("Could not create perfdata: %v", err)
	}
	pd = append(pd, *datum)

	// 'label'=value[UOM];[warn];[crit];[min];[max]
	expected := " | age=0.123s;;0.5;0;"
	result := RenderPerfdata(pd)
	if result != expected {
		t.Errorf("Perfdata rendering error: expected %s, got %v", expected, result)
	}
}
