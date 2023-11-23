package leetcode

// @Title        0852.Peak-Index-In-A-Mountain-Array.go
// @Description  0852.Peak-Index-In-A-Mountain-Array solution
// @Create       XdpCs 2023-11-23 11:11
// @Update       XdpCs 2023-11-23 11:11

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	ans := 0
	for left <= right {
		mid := (left + right) >> 1
		if arr[mid] > arr[mid+1] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
