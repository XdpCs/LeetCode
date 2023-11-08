# [0258. 各位相加](https://leetcode.cn/problems/add-digits/)

## 题目描述

给定一个非负整数`num`，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。

**示例 1：**

```text
输入: num = 38
输出: 2 
解释: 各位相加的过程为：
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
由于 2 是一位数，所以返回 2。
```

**示例 2：**

```text
输入: num = 0
输出: 0
```

**限制**

* 0 <= num <= 2<sup>31</sup> - 1

## 题目大意

给定一个非负整数`num`，反复将各个位上的数字相加，直到结果为一位数。返回这个结果

## 解题思路

时间复杂度：O(log num)，对于计算一次各位相加需要 O(log num)的时间，由于num <= 2<sup>31</sup> -
1，所以计算一次各位相加最大可能结果是82

空间复杂度：O(1)

这个题目，按照题意，直接循环即可

## 代码

```go
package leetcode

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for ; num > 0; num /= 10 {
			sum += num % 10
		}
		num = sum
	}
	return num
}
```