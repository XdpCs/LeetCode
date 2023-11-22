package leetcode

// @Title        0263.Ugly-Number.go
// @Description  0263.Ugly-Number solution
// @Create       XdpCs 2023-11-22 20:24
// @Update       XdpCs 2023-11-22 20:24

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	factor := []int{2, 3, 5}
	for _, f := range factor {
		for n%f == 0 {
			n = n / f
		}
	}
	return n == 1
}
