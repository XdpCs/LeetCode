package main

// @Title        0326.Power-Of-Three.go
// @Description  0326.Power-Of-Three solution
// @Create       XdpCs 2023-11-13 11:16
// @Update       XdpCs 2023-11-13 11:16

import "math"

func isPowerOfThree(n int) bool {
	hashSet := map[int]struct{}{}
	for i := 1; i <= math.MaxInt32; i *= 3 {
		hashSet[i] = struct{}{}
	}
	_, ok := hashSet[n]
	return ok
}

func isPowerOfThreeMath(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func GetMaxInt32PowerOfThree() int {
	i := 1
	for i <= math.MaxInt32 {
		i *= 3
	}
	return i / 3
}
