# [1422.分割字符串的最大得分](https://leetcode.cn/problems/maximum-score-after-splitting-a-string)

## 题目描述

给你一个由若干`0`和`1`组成的字符串`s`，请你计算并返回将该字符串分割成两个`非空`子字符串（即`左`子字符串和`右`
子字符串）所能获得的最大得分。

「分割字符串的得分」为`左`子字符串中`0`的数量加上`右`子字符串中`1`的数量。

**示例 1：**

> 输入：s = "011101"
> 输出：5
> 解释：
> 将字符串 s 划分为两个非空子字符串的可行方案有：
> 左子字符串 = "0" 且 右子字符串 = "11101"，得分 = 1 + 4 = 5
> 左子字符串 = "01" 且 右子字符串 = "1101"，得分 = 1 + 3 = 4
> 左子字符串 = "011" 且 右子字符串 = "101"，得分 = 1 + 2 = 3
> 左子字符串 = "0111" 且 右子字符串 = "01"，得分 = 1 + 1 = 2
> 左子字符串 = "01110" 且 右子字符串 = "1"，得分 = 2 + 1 = 3

**示例 2：**

> 输入：s = "00111"
> 输出：5
> 解释：当 左子字符串 = "00" 且 右子字符串 = "111" 时，我们得到最大得分 = 2 + 3 = 5

**示例 3：**

> 输入：s = "1111"
> 输出：3

**限制**

* 2 <= s.length <= 500
* 字符串 s 仅由字符`'0'`和`'1'`组成。

## 题目大意

左子串中`0`的数量加上右子串中`1`的数量被称为一个「得分」。对`s`进行分割，使得得分最大 。

## 解题思路

时间复杂度：O(n)

空间复杂度：O(1)

这道题我们可以发现整个字符串的`0`和`1`
是固定的，我们可以通过分割字符串，将字符串分为左右两个部分，然后分别统计左右两个部分的`0`和`1`
的数量，然后将两个数量相加，得到最大的得分。我们可以发现如果分割点右移动，如果该数字为`0`，那么左边的`0`加一，如果该数字为`1`
，那么右边的`1`减一，我们可以通过这种方式来计算出每个分割点的得分，然后取最大值即可。

## 代码

```go
package leetcode

func maxScore(s string) int {
	leftScore, rightScore := 0, 0
	if judgeLeft(s[0]) {
		leftScore++
	}
	for i := 1; i < len(s); i++ {
		if judgeRight(s[i]) {
			rightScore++
		}
	}
	maxScore := leftScore + rightScore

	for i := 1; i < len(s)-1; i++ {
		if judgeLeft(s[i]) {
			leftScore++
		} else {
			rightScore--
		}
		maxScore = max(maxScore, leftScore+rightScore)
	}

	return maxScore
}

func judgeLeft(n byte) bool {
	return judgeNum(n, '0')
}

func judgeRight(n byte) bool {
	return judgeNum(n, '1')
}

func judgeNum(n byte, num byte) bool {
	if n == num {
		return true
	}
	return false
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
```
