package leetcode

// @Title        0231.Power-Of-Two.go
// @Description  0231.Power-Of-Two solution
// @Create       XdpCs 2023-11-13 11:16
// @Update       XdpCs 2023-11-13 11:16

import "math"

func isPowerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}

func isPowerOfTwoStorage(n int) bool {
	hashSet := map[int]struct{}{}
	for i := 1; i <= math.MaxInt32; i *= 2 {
		hashSet[i] = struct{}{}
	}
	_, ok := hashSet[n]
	return ok
}

func isPowerOfTwoMaxNum(n int) bool {
	maxPowerOfTwo := 1 << 30
	return (n > 0) && (maxPowerOfTwo%n == 0)
}

func isPowerOfTwoLoop(n int) bool {
	ans := 1
	for ; n > ans; ans = ans << 1 {
	}
	if n == ans {
		return true
	}
	return false
}
