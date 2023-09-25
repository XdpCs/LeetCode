# [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/)

## 题目描述

给你一个整数数组 `nums` 。

如果一组数字 `(i,j)` 满足 `nums[i]` == `nums[j]` 且 `i < j`，就可以认为这是一组**好数对**。

返回好数对的数目。

**示例 1：**

```
输入：nums = [1,2,3,1,1,3]
输出：4
解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始
```

**示例 2：**

```
输入：nums = [1,1,1,1]
输出：6
解释：数组中的每组数字都是好数对
```

**示例 3：**

```
输入：nums = [1,2,3]
输出：0
```

**限制**

* `1 <= nums.length <= 100`
* `1 <= nums[i] <= 100`

## 题目大意

给你一个整数数组 `nums` ，如果一组数字 `(i,j)` 满足 `nums[i]` == `nums[j]` 且 `i < j`，就可以认为这是一组**好数对**。

## 解题思路

时间复杂度：O(n)

空间复杂度：O(n)

先遍历输入的整数数组`nums`，然后用`哈希表`存储每个数字的数量，最后遍历`哈希表`
，判断是否大于`1`，计算每个数字可以组成的好数对的数量：这是一个组合问题，组合的公式是$\frac{m!}{n!*(m-n)!}$。`m`
是每个数字的数量，`n`是`2`，通过约分可得到这个公式 $\frac{m*(m-1)}{2}$，即为这个数组成的好数对的数量，然后将每个数的好数对数量相加即为结果。

## 代码

```go
package leetcode

func numIdenticalPairs(nums []int) int {
	n := len(nums)
	hashMap := make(map[int]int, n)
	sum := 0
	for _, num := range nums {
		hashMap[num]++
	}

	for _, v := range hashMap {
		if v > 1 {
			sum += v * (v - 1) / 2
		}
	}
	return sum
}
```
