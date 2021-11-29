package easy

import (
	"testing"
)

func TestTournamentWinner(t *testing.T) {
	testComps := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	testResults := []int{0, 0, 1}
	result := TournamentWinner(testComps, testResults)
	if result != "Python" {
		t.Errorf("Expected Python, got %s", result)
	}
}
