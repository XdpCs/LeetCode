# [0852.山脉数组的峰顶索引](https://leetcode.cn/problems/peak-index-in-a-mountain-array/)

## 题目描述

符合下列属性的数组 arr 称为 山脉数组 ：

* arr.length >= 3
* 存在 i（0 < i < arr.length - 1）使得：
  * arr[0] < arr[1] < ... arr[i-1] < arr[i]
  * arr[i] > arr[i+1] > ... > arr[arr.length - 1]

给你由整数组成的山脉数组`arr`，返回满足`arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... >
arr[arr.length - 1]`的下标 i 。

你必须设计并实现时间复杂度为`O(log(n))`的解决方案。

### 示例 1

> 输入：arr = [0,1,0]
> 输出：1

### 示例 2

> 输入：arr = [0,2,1,0]
> 输出：1

### 示例 3

> 输入：arr = [0,10,5,2]
> 输出：1

### 限制

* 3 <= arr.length <= $10^5$
* 0 <= arr[i] <= $10^6$
* 题目数据保证`arr`是一个山脉数组

## 题目大意

找到数组中最大的元素，输出它的下标。题目保证数组中最大的元素是唯一的。

## 解题思路

时间复杂度：O(log n)

空间复杂度：O(1)

题目中强制需要时间复杂度为 O(log n) 的解法，再根据已经给出了数组的性质，数组中最大的元素是唯一的，且最大元素左边的元素是递增的，最大元素右边的元素是递减的。因此可以使用二分查找的方法，找到最大元素的下标。

## 代码

```go
package leetcode

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
```
