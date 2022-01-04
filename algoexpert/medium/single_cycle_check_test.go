package medium

import "testing"

func TestHasSingleCycle(t *testing.T) {
	noCycle := []int{0, 1, 1, 1, 1}
	cycle := []int{2, 3, 1, -4, -4, 2}

	if HasSingleCycle(noCycle) {
		t.Error("Should not have cycle")
	}

	if HasSingleCycle(cycle) == false {
		t.Error("Should have cycle")
	}
}
