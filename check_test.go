package nagiosplugin

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	c := NewCheck()
	expected := "CRITICAL: 200000 terrifying space monkeys in the engineroom | space_monkeys=200000c;10000;100000;0;4294967296"
	nSpaceMonkeys := float64(200000)
	maxSpaceMonkeys := float64(1 << 32)
	c.AddPerfDatum("space_monkeys", "c", nSpaceMonkeys, 0, maxSpaceMonkeys, 10000, 100000)
	c.AddResult(CRITICAL, fmt.Sprintf("%v terrifying space monkeys in the engineroom", nSpaceMonkeys))
	// Check a WARNING can't override a CRITICAL
	c.AddResult(WARNING, fmt.Sprintf("%v slightly annoying space monkeys in the engineroom", nSpaceMonkeys))
	result := c.String()
	if expected != result {
		t.Errorf("Expected check output %v, got check output %v", expected, result)
	}
}

func TestDefaultStatusPolicy(t *testing.T) {
	c := NewCheck()
	c.AddResult(WARNING, "Isolated-frame flux emission outside threshold")
	c.AddResult(UNKNOWN, "No response from betaform amplifier")

	expected := "UNKNOWN"
	actual := strings.SplitN(c.String(), ":", 2)[0]
	if actual != expected {
		t.Errorf("Expected %v status, got %v", expected, actual)
	}
}

func TestCustomStatusPolicy(t *testing.T) {
	p, _ := NewStatusPolicy([]Status{OK, UNKNOWN, WARNING, CRITICAL})
	c := NewCheckWithOptions(CheckOptions{
		StatusPolicy: p,
	})
	c.AddResult(WARNING, "Isolated-frame flux emission outside threshold")
	c.AddResult(UNKNOWN, "No response from betaform amplifier")

	expected := "WARNING"
	actual := strings.SplitN(c.String(), ":", 2)[0]
	if actual != expected {
		t.Errorf("Expected %v status, got %v", expected, actual)
	}
}
