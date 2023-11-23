package leetcode

// @Title        1422.Maximum-Score-After-Splitting-A-String.go
// @Description	 1422.Maximum-Score-After-Splitting-A-String solution
// @Create       XdpCs 2023-11-23 10:30
// @Update       XdpCs 2023-11-23 10:30

func maxScore(s string) int {
	leftScore, rightScore := 0, 0
	if judgeLeft(s[0]) {
		leftScore++
	}
	for i := 1; i < len(s); i++ {
		if judgeRight(s[i]) {
			rightScore++
		}
	}
	maxScore := leftScore + rightScore

	for i := 1; i < len(s)-1; i++ {
		if judgeLeft(s[i]) {
			leftScore++
		} else {
			rightScore--
		}
		maxScore = max(maxScore, leftScore+rightScore)
	}

	return maxScore
}

func judgeLeft(n byte) bool {
	return judgeNum(n, '0')
}

func judgeRight(n byte) bool {
	return judgeNum(n, '1')
}

func judgeNum(n byte, num byte) bool {
	if n == num {
		return true
	}
	return false
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
