package medium

import "testing"

func TestMinimumNumberOfCoins(t *testing.T) {
	result := MinimumNumberOfCoins(6, []int{1, 2, 4})
	if result != 2 {
		t.Errorf("Unexpected amount: %d", result)
	}
}
