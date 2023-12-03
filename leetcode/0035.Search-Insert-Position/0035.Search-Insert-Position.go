package leetcode

// @Title        0035.Search-Insert-Position.go
// @Description
// @Create       XdpCs 2023-12-03 19:06
// @Update       XdpCs 2023-12-03 19:06

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := len(nums)
	for left <= right {
		mid := (right-left)>>1 + left
		ans := nums[mid]
		if ans >= target {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}
