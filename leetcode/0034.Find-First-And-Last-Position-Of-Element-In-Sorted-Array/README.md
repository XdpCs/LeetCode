# [0034.在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

## 题目描述

给你一个按照非递减顺序排列的整数数组`nums`，和一个目标值`target`。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值`target`，返回 [-1, -1]。

你必须设计并实现时间复杂度为`O(log n)`的算法解决此问题。

**示例 1：**

> 输入：nums = [5,7,7,8,8,10], target = 8
> 输出：[3,4]

**示例 2：**

> 输入：nums = [5,7,7,8,8,10], target = 6
> 输出：[-1,-1]

**示例 3：**

> 输入：nums = [], target = 0
> 输出：[-1,-1]

**限制**

* 0 <= nums.length <= $10^5$
* $-10^9$ <= nums[i] <= $10^9$
* `nums`是一个非递减数组
* $-10^9$ <= target <= $10^9$

## 题目大意

给定一个按照非递减顺序排列的整数数组`nums`，和一个目标值`target`。请你找出给定目标值在数组中的开始位置和结束位置。

## 解题思路

时间复杂度：O(log n)

空间复杂度：O(1)

这个问题可以拆分成两个小问题：[查找第一个等于目标值的元素](../../notes/Binary-Search.md#查找第一个等于目标值的元素)
和[查找最后一个等于目标值的元素](../../notes/Binary-Search.md#查找最后一个等于目标值的元素)。

### 代码

```go
package leetcode

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
```