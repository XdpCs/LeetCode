# [0001.两数之和](https://leetcode.cn/problems/two-sum/)

## 题目描述

给定一个整数数组`nums`和一个整数目标值`target`，请你在该数组中找出和为目标值`target` 的那**两个**整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

### 示例 1

> 输入：nums = [2,7,11,15], target = 9
> 输出：[0,1]
> 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

### 示例 2

> 输入：nums = [3,2,4], target = 6
> 输出：[1,2]

### 示例 3

> 输入：nums = [3,3], target = 6
> 输出：[0,1]

### 限制

* 2 <= nums.length <= $10^4$
* $-10^9$ <= nums[i] <= $10^9$
* $-10^9$ <= target <= $10^9$
* 只会存在一个有效答案

## 题目大意

找出两个下标并且下标对应的值相加等于目标值

## 解题思路

### 方法一 (循环)

时间复杂度：O($n^2$)

空间复杂度：O(1)

直接循环遍历数组，找出两个数相加等于目标值的下标

### 代码

```go
package leetcode

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			} else if nums[i]+nums[j] > target {
				continue
			}
		}
	}
	return []int{0, 0}
}
```

### 方法二 (哈希)

时间复杂度：O(n)

空间复杂度：O(n)

使用哈希表存储数组中的值和下标，遍历数组，判断目标值减去当前值是否在哈希表中，如果在则返回当前下标和哈希表中的下标

### 代码

```go
package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range nums {
		if j, ok := hashMap[target-v]; ok {
			return []int{i, j}
		}
		hashMap[v] = i
	}
	return []int{0, 0}
}
```
