package leetcode

// @Title        0704.Binary-Search.go
// @Description
// @Create       XdpCs 2023-12-03 18:01
// @Update       XdpCs 2023-12-03 18:01

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := (right-left)>>1 + left
		ans := nums[mid]
		if ans == target {
			result = mid
			right = mid - 1
		} else if ans > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}
