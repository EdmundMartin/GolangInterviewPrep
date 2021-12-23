package medium

import (
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	heap := NewMinHeap([]int{})
	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(30)
	heap.Insert(40)
	heap.Insert(24)

	for _, val := range []int{10, 20, 24, 30, 40} {
		popped := heap.Remove()
		if popped != val {
			t.Errorf("Expected val: %d got: %d", val, popped)
		}
	}
}
