package leetcode

// @Title        0069.Sqrtx.go
// @Description
// @Create       XdpCs 2023-12-04 16:58
// @Update       XdpCs 2023-12-04 16:58

func mySqrt(x int) int {
	l, r := 0, x
	result := -1
	for l <= r {
		mid := (r-l)>>1 + l
		if mid*mid <= x {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}
