package leetcode

// @Title        0034.Find-First-And-Last-Position-Of-Element-In-Sorted-Array.go
// @Description
// @Create       XdpCs 2023-12-03 19:20
// @Update       XdpCs 2023-12-03 19:20

func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums, target), searchLast(nums, target)}
}

func searchFirst(nums []int, target int) int {
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

func searchLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	for left <= right {
		mid := (right-left)>>1 + left
		ans := nums[mid]
		if ans < target {
			left = mid + 1
		} else if ans == target {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
