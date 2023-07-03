package main

import (
	"testing"
)

func TestValidateRules(t *testing.T) {
	rules := []rule{
		{
			"LL100",
			"Generic error",
			[]string{"generic"},
			"Errors are bad.",
			"(?i).*error:.*",
		},
	}
	expectedExitStatus := 0
	exitStatus := validateRules(rules, "")
	if exitStatus != expectedExitStatus {
		t.Fatalf("Expected exit status '%d', got '%d'", exitStatus, expectedExitStatus)
	}
}
