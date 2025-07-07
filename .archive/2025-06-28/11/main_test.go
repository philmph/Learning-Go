package main

import (
	"testing"
)

func TestCalculateSomething(t *testing.T) {
	// Log information (only shows with -v flag or if test fails)
	t.Log("This is a log message")

	// Test a condition
	x, y := 5, 5
	result := calculateSomething(x, y)
	expected := 25

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result) // Report error but continue
	}

	t.Logf("Passed")
}
