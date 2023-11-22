# [0867.转置矩阵](https://leetcode.cn/problems/transpose-matrix/)

## 题目描述

给你一个二维整数数组`matrix`， 返回`matrix`的**转置矩阵**。

矩阵的`转置`是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

![example](./images/example.png)

**示例 1：**

> 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
> 输出：[[1,4,7],[2,5,8],[3,6,9]]

**示例 2：**

> 输入：matrix = [[1,2,3],[4,5,6]]
> 输出：[[1,4],[2,5],[3,6]]

**限制**

* m == matrix.length
* n == matrix[i].length
* 1 <= m, n <= 1000
* 1 <= m * n <= $10^5$
* $-10^9$ <= matrix[i][j] <= $10^9$

## 题目大意

给你一个二维整数数组 `matrix`， 返回 `matrix` 的 **转置矩阵** 。

## 解题思路

时间复杂度：O(m*n)m, n 分别为矩阵的行数和列数

空间复杂度：O(1), 返回值不计入空间复杂度

这个题目，按照题意，直接循环即可

## 代码

```go
package leetcode

func transpose(matrix [][]int) [][]int {
	ansMatrix := make([][]int, len(matrix[0]))
	for i := 0; i < len(ansMatrix); i++ {
		ansMatrix[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ansMatrix[j][i] = matrix[i][j]
		}
	}

	return ansMatrix
}
```
