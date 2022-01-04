package medium

func SearchInSortedMatrix(matrix [][]int, target int) []int {
	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		currentValue := matrix[row][col]
		if currentValue > target {
			col -= 1
		} else if currentValue < target {
			row += 1
		} else {
			return []int{row, col}
		}
	}
	return []int{-1, -1}
}
