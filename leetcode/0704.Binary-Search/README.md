# [0704.二分查找](https://leetcode.cn/problems/binary-search/)

## 题目描述

给定一个`n`个元素有序的`（升序）`整型数组`nums`和一个目标值`target`，写一个函数搜索`nums`中的`target`，如果目标值存在返回下标，否则返回
`-1`。

**示例 1：**

> 输入: nums = [-1,0,3,5,9,12], target = 9
> 输出: 4
> 解释: 9 出现在 nums 中并且下标为 4

**示例 2：**

> 输入: nums = [-1,0,3,5,9,12], target = 2
> 输出: -1
> 解释: 2 不存在 nums 中因此返回 -1

**限制**

* 你可以假设`nums`中的所有元素是不重复的。
* `n`将在 [1, 10000]之间。
* `nums`的每个元素都将在 [-9999, 9999]之间。

## 题目大意

数组是有序的，可以使用二分查找的方法。找到目标值的下标。

## 解题思路

时间复杂度：O(log n)

空间复杂度：O(1)

这其实二分查找中[查找第一个等于目标值的元素](../../notes/Binary-Search.md#查找第一个等于目标值的元素)的变形

### 代码

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