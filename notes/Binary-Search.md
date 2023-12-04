# 二分查找

## 查找第一个等于目标值的元素

```go
package leetcode

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
```

## 查找最后一个等于目标值的元素

```go
package leetcode

func search(nums []int, target int) int {
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
```

## 查找第一个大于等于目标值的元素

```go
package leetcode

func search(nums []int, target int) int {
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
```

## 查找最后一个小于等于目标值的元素

```go
package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := len(nums)
	for left <= right {
		mid := (right-left)>>1 + left
		ans := nums[mid]
		if ans <= target {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
```
