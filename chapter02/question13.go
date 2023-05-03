package chapter02

type Matrix struct {
	matrix [][]int
	sum    [][]int
}

func newMatrix(matrix [][]int) *Matrix {
	sum := make([][]int, 0, len(matrix))

	for i, row := range matrix {
		curr, rowSum := 0, make([]int, len(row))
		for j, col := range row {
			curr += col
			rowSum[j] = curr
			if i > 0 {
				rowSum[j] += sum[i-1][j]
			}
		}
		sum = append(sum, rowSum)
	}

	return &Matrix{matrix: matrix, sum: sum}
}

func (m *Matrix) calculate(x1, y1, x2, y2 int) int {
	minX, minY, maxX, maxY := 0, 0, 0, 0

	if x1 < x2 {
		minX, maxX = x1-1, x2
	} else {
		minX, maxX = x2-1, x1
	}

	if y1 < y2 {
		minY, maxY = y1-1, y2
	} else {
		minY, maxY = y2-1, y1
	}

	return m.sum[maxX][maxY] - m.sum[maxX][minY] - m.sum[minX][maxY] + m.sum[minX][minY]
}
