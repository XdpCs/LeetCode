package leetcode

// @Title        0367.Valid-Perfect-Square.go
// @Description
// @Create       XdpCs 2023-12-04 17:18
// @Update       XdpCs 2023-12-04 17:18

func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l <= r {
		m := (r-l)>>1 + l
		ans := m * m
		if ans == num {
			return true
		} else if ans < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
