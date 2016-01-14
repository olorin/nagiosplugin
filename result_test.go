package nagiosplugin

import (
	"testing"
)

func TestNewStatusPolicyAcceptsCompleteStatuses(t *testing.T) {
	_, err := NewStatusPolicy([]Status{OK, UNKNOWN, WARNING, CRITICAL})
	if err != nil {
		t.Errorf("NewStatusPolicy(): %v", err)
	}
}

func TestNewStatusPolicyRejectsIncompleteStatuses(t *testing.T) {
	// Missing UNKNOWN.
	_, err := NewStatusPolicy([]Status{OK, WARNING, CRITICAL})
	if err == nil {
		t.Errorf("expected NewStatusPolicy() to return an error")
	}
}
