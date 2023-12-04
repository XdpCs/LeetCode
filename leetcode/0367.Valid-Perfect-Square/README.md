# [0367.有效的完全平方数](https://leetcode.cn/problems/valid-perfect-square/)

## 题目描述

给你一个正整数`num`。如果`num`是一个完全平方数，则返回`true`，否则返回`false`。

**完全平方数**是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

不能使用任何内置的库函数，如 `sqrt`。

**示例 1：**

> 输入：num = 16 
> 输出：true 
> 解释：返回 true ，因为 4 * 4 = 16 且 4 是一个整数

**示例 2：**

> 输入：num = 14 
> 输出：false 
> 解释：返回 false ，因为 3.742 * 3.742 = 14 但 3.742 不是一个整数。

**限制**

* 1 <= n <= $2^{31}$ - 1

## 题目大意

判断`num`是不是某一个整数的完全平方数

## 解题思路


时间复杂度：O(log n)

空间复杂度：O(1)

使用二分查找，找到第一个等于目标值的元素的数，即为`num`的完全平方数。这是[查找第一个等于目标值的元素](../../notes/Binary-Search.md#查找第一个等于目标值的元素)的变形。

## 代码

```go
package leetcode

func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l <= r {
		m := (r-l)>>1 + l
		ans := m * m
		if ans == num {
			return true
		} else if ans < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
```

