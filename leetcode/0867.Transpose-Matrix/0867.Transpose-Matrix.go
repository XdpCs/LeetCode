package leetcode

// @Title        0867.Transpose-Matrix.go
// @Description  0867.Transpose-Matrix solution
// @Create       XdpCs 2023-11-22 20:57
// @Update       XdpCs 2023-11-22 20:57

func transpose(matrix [][]int) [][]int {
	ansMatrix := make([][]int, len(matrix[0]))
	for i := 0; i < len(ansMatrix); i++ {
		ansMatrix[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ansMatrix[j][i] = matrix[i][j]
		}
	}

	return ansMatrix
}
