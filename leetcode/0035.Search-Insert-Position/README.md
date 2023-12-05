# [0035.搜索插入位置](https://leetcode.cn/problems/search-insert-position/)

## 题目描述

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为`O(log n)`的算法。

### 示例 1

> 输入: nums = [1,3,5,6], target = 5
> 输出: 2

### 示例 2

> 输入: nums = [1,3,5,6], target = 2
> 输出: 1

### 示例 3

> 输入: nums = [1,3,5,6], target = 7
> 输出: 4

### 限制

* 1 <= nums.length <= $10^4$
* $-10^4$ <= nums[i] <= $10^4$
* `nums`为`无重复元素`的`升序`排列数组
* $-10^4$ <= target <= $10^4$

## 题目大意

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

## 解题思路

时间复杂度：O(log n)

空间复杂度：O(1)

这其实二分查找中[查找第一个大于等于目标值的元素](../../notes/Binary-Search.md#查找第一个大于等于目标值的元素)
的变形，需要注意的是，如果是最后一个元素都小于目标值，那么返回的是`len(nums)`，所以开始的时候`result`的初始值为`len(nums)`。

### 代码

```go
package leetcode

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
```
