package leetcode

// @Title        2236.Root-Equals-Sum-of-Children.go
// @Description  2413.Smallest-Even-Multiple solution
// @Create       XdpCs 2023-09-19 17:31
// @Update       XdpCs 2023-09-19 17:31

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	} else {
		return 2 * n
	}
}
