package leetcode

// @Title        1470.Shuffle-The-Array.go
// @Description  1470.Shuffle-The-Array solution
// @Create       XdpCs 2023-11-22 20:37
// @Update       XdpCs 2023-11-22 20:37

func shuffle(nums []int, n int) []int {
	ans := make([]int, 0, 2*n)
	for i, j := 0, n; i < n; i++ {
		ans = append(ans, nums[i])
		ans = append(ans, nums[j])
		j++
	}
	return ans
}
