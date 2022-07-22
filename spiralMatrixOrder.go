func spiralMatrix(matrix [][]int) []int {

	var res []int
	top := 0
	right := len(matrix[0]) - 1
	down := len(matrix) - 1
	left := 0
	dir := 0

	for top <= down && left <= right {
		if dir == 0 {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		}
		if dir == 1 {
			for i := top; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}
		if dir == 2 {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			down--
		}
		if dir == 3 {
			for i := down; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
		dir = (dir + 1) % 4
	}

	return res
}
