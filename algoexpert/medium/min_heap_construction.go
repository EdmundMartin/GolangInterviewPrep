package medium

type MinHeap struct {
	heap []int
}

func NewMinHeap(values []int) *MinHeap {
	m := &MinHeap{heap: values}
	firstParentIdx := (len(m.heap) - 2) / 2
	for idx := firstParentIdx; idx >= 0; idx-- {
		m.heap = m.siftDown(idx, len(m.heap), m.heap)
	}
	return m
}

func (m *MinHeap) siftDown(currentIdx, endIdx int, heap []int) []int {
	childOneIdx := currentIdx*2 + 1
	for childOneIdx <= endIdx {
		childTwoIdx := currentIdx*2 + 2
		if childTwoIdx > endIdx {
			childTwoIdx = -1
		}
		var idxSwap int
		if childTwoIdx != -1 && heap[childTwoIdx] < heap[childOneIdx] {
			idxSwap = childTwoIdx
		} else {
			idxSwap = childOneIdx
		}
		if heap[idxSwap] < heap[currentIdx] {
			heap[idxSwap], heap[currentIdx] = heap[currentIdx], heap[idxSwap]
			currentIdx = idxSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			break
		}
	}
	return heap
}

func (m *MinHeap) siftUp(currentIdx int, heap []int) []int {
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && heap[currentIdx] < heap[parentIdx] {
		heap[currentIdx], heap[parentIdx] = heap[parentIdx], heap[currentIdx]
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
	return heap
}

func (m *MinHeap) Peek() int {
	return m.heap[0]
}

func (m *MinHeap) Insert(value int) {
	m.heap = append(m.heap, value)
	m.heap = m.siftUp(len(m.heap)-1, m.heap)
}

func (m *MinHeap) Remove() int {
	m.heap[0], m.heap[len(m.heap)-1] = m.heap[len(m.heap)-1], m.heap[0]
	var val int
	val, m.heap = m.heap[len(m.heap)-1], m.heap[:len(m.heap)-1]
	m.heap = m.siftDown(0, len(m.heap)-1, m.heap)
	return val
}
