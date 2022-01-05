package medium

type MinBst struct {
	Value int
	Left *MinBst
	Right *MinBst
}


func (m *MinBst) insert(value int) {
	if value < m.Value {
		if m.Left == nil {
			m.Left = &MinBst{Value: value}
		} else {
			m.Left.insert(value)
		}
	} else {
		if m.Right == nil {
			m.Right = &MinBst{Value: value}
		} else {
			m.Right.insert(value)
		}
	}
}


func constructMinHeightBst(array []int, startIdx, endIdx int) *MinBst {
	if endIdx < startIdx {
		return nil
	}
	middleIdx := (startIdx + endIdx) / 2
	bst := &MinBst{Value: array[middleIdx]}
	bst.Left = constructMinHeightBst(array, startIdx, middleIdx - 1)
	bst.Right = constructMinHeightBst(array, middleIdx + 1, endIdx)
	return bst
}


func MinHeightBST(array []int) *MinBst {
	return constructMinHeightBst(array, 0, len(array) - 1)
}