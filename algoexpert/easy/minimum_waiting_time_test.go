package easy

import "testing"

func TestMinimumWaitingTime(t *testing.T) {
	testData := []int{3, 2, 1, 2, 6}
	result := MinimumWaitingTime(testData)
	if result != 17 {
		t.Errorf("Expected 17, got %d", result)
	}
}