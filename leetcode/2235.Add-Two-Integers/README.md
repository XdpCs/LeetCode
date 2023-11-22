# [2235.两整数相加](https://leetcode.cn/problems/add-two-integers/)

## 题目描述

给你两个整数`num1`和`num2`，返回这两个整数的和。

**示例 1：**

> 输入：num1 = 12, num2 = 5
> 输出：17
> 解释：num1 是 12，num2 是 5 ，它们的和是 12 + 5 = 17 ，因此返回 17 。

**示例 2：**

> 输入：num1 = -10, num2 = 4
> 输出：-6
> 解释：num1 + num2 = -6 ，因此返回 -6 。

**限制**

* `-100 <= num1, num2 <= 100 `

## 题目大意

给你两个整数`num1`和`num2`，返回这两个整数的和。

## 解题思路

* 时间复杂度：O(1)
* 空间复杂度：O(1)
* 通过加法运算符 `+`，将两个整数 `num1` 和 `num2` 相加。

## 代码

```go
package leetcode

func sum(num1 int, num2 int) int {
	return num1 + num2
}
```
